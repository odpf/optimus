package extension

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/google/go-github/github"
)

const defaultRepoPrefix = "optimus-extension"

// Extension is manager for extension
type Extension struct {
	ctx context.Context

	githubClient *github.Client
	httpClient   *http.Client

	manifest *Manifest

	reservedCommands []string
}

// NewExtension initializes extension
func NewExtension(
	ctx context.Context, manifest *Manifest,
	githubClient *github.Client, httpClient *http.Client,
	reservedCommands ...string,
) (*Extension, error) {
	return &Extension{
		ctx:              ctx,
		githubClient:     githubClient,
		httpClient:       httpClient,
		manifest:         manifest,
		reservedCommands: reservedCommands,
	}, nil
}

// Install installs extension from a Gilthub repository
func (e *Extension) Install(owner, repo, alias string) error {
	if err := e.validateInstall(owner, repo, alias); err != nil {
		return err
	}
	release, _, err := e.githubClient.Repositories.GetLatestRelease(e.ctx, owner, repo)
	if err != nil {
		return fmt.Errorf("error getting the latest release: %v", err)
	}
	downloadURL, err := e.getDownloadURL(release)
	if err != nil {
		return err
	}
	extensionDir, err := getExtensionDir()
	if err != nil {
		return err
	}
	destDirPath := path.Join(extensionDir, owner, repo)
	if err := os.MkdirAll(destDirPath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating dir: %v", err)
	}
	tag := repo
	if release.TagName != nil {
		tag = *release.TagName
	}
	destFilePath := path.Join(destDirPath, tag)
	if err := e.downloadAsset(downloadURL, destFilePath); err != nil {
		return err
	}
	name := strings.TrimPrefix(repo, defaultRepoPrefix)
	aliases := []string{name}
	if alias != "" {
		aliases = append(aliases, alias)
	}
	metadata := &Metadata{
		Owner:     owner,
		Repo:      repo,
		Aliases:   aliases,
		Tag:       tag,
		LocalPath: destFilePath,
	}
	e.manifest.Metadatas = append(e.manifest.Metadatas, metadata)
	e.manifest.Update = time.Now()
	return FlushManifest(e.manifest)
}

func (e *Extension) downloadAsset(url, destPath string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Accept", "application/octet-stream")

	resp, err := e.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return e.getResponseError(resp)
	}

	f, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}
	return nil
}

func (e *Extension) getResponseError(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %v", err)
	}
	var buff bytes.Buffer
	if err := json.Indent(&buff, body, "", "  "); err != nil {
		return fmt.Errorf("error indenting json: %v", err)
	}
	return errors.New(string(buff.Bytes()))
}

func (e *Extension) getDownloadURL(release *github.RepositoryRelease) (string, error) {
	currentDist := e.getCurrentDist()
	for _, asset := range release.Assets {
		if asset.Name != nil {
			if strings.HasSuffix(*asset.Name, currentDist) && asset.BrowserDownloadURL != nil {
				return asset.GetBrowserDownloadURL(), nil
			}
		}
	}
	return "", fmt.Errorf("asset for [%s] is not found", currentDist)
}

func (e *Extension) getCurrentDist() string {
	return runtime.GOOS + "-" + runtime.GOARCH
}

func (e *Extension) validateInstall(owner, repo, alias string) error {
	if !strings.HasPrefix(repo, defaultRepoPrefix) {
		return fmt.Errorf("[%s] does not have prefix [%s]", repo, defaultRepoPrefix)
	}
	name := strings.TrimPrefix(repo, defaultRepoPrefix)
	for _, c := range e.reservedCommands {
		if name == c {
			return fmt.Errorf("[%s] is reserved command", name)
		}
	}
	for _, metadata := range e.manifest.Metadatas {
		if owner == metadata.Owner && repo == metadata.Repo {
			return fmt.Errorf("%s/%s [%s] is already installed", owner, repo, metadata.Tag)
		}
		for _, a := range metadata.Aliases {
			if alias == a {
				return fmt.Errorf("alias [%s] is already used", alias)
			}
		}
	}
	return nil
}

// Run executes extension
func (e *Extension) Run(name string, args []string) error {
	var path string
	for _, metadata := range e.manifest.Metadatas {
		for _, a := range metadata.Aliases {
			if name == a {
				path = metadata.LocalPath
			}
		}
	}
	cmd := exec.Command(path, args...)
	b, err := cmd.Output()
	fmt.Println(string(b))
	return err
}
