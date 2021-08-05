package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-version"

	"github.com/odpf/optimus/config"
	"github.com/odpf/optimus/models"

	pb "github.com/odpf/optimus/api/proto/odpf/optimus"
	"github.com/pkg/errors"
	cli "github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	versionTimeout   = time.Second * 2
	clientVersion    = time.Second * 1
	githubReleaseURL = "https://api.github.com/repos/odpf/optimus/releases/latest"
)

func versionCommand(l logger, host string, pluginRepo models.PluginRepository) *cli.Command {
	var serverVersion bool
	c := &cli.Command{
		Use:   "version",
		Short: "Print the client version information",
		RunE: func(c *cli.Command, args []string) error {
			l.Printf(fmt.Sprintf("client: %s-%s", coloredNotice(config.Version), config.BuildCommit))
			if host != "" && serverVersion {
				srvVer, err := getVersionRequest(config.Version, host)
				if err != nil {
					return err
				}
				l.Printf("server: %s", coloredNotice(srvVer))
			}
			checkLatestVersion(l)

			plugins := pluginRepo.GetAll()
			l.Println(fmt.Sprintf("\nDiscovered plugins: %d", len(plugins)))
			for taskIdx, tasks := range plugins {
				schema := tasks.Info()
				l.Printf("%d. %s\n", taskIdx+1, schema.Name)
				l.Printf("Description: %s\n", schema.Description)
				l.Printf("Image: %s\n", schema.Image)
				l.Printf("Type: %s\n", schema.PluginType)
				l.Printf("Plugin version: %s\n", schema.PluginVersion)
				l.Printf("Plugin mods: %v\n", schema.PluginMods)
				if schema.HookType != "" {
					l.Printf("Hook type: %s\n", schema.HookType)
				}
				if len(schema.DependsOn) != 0 {
					l.Printf("Depends on: %v\n", schema.DependsOn)
				}
				l.Println("")
			}
			return nil
		},
	}
	c.Flags().BoolVar(&serverVersion, "with-server", false, "check for server version")
	return c
}

func checkLatestVersion(l logger) {
	gitClient := http.Client{
		Timeout: clientVersion,
	}

	req, err := http.NewRequest(http.MethodGet, githubReleaseURL, nil)
	if err != nil {
		l.Println("failed to create request for latest version")
		return
	}
	req.Header.Set("User-Agent", "optimus")
	res, err := gitClient.Do(req)
	if err != nil {
		l.Println("failed to get latest version from github")
		return
	}
	if res.StatusCode != http.StatusOK {
		l.Println("failed to get latest version from github")
		return
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l.Println("failed to read response body")
		return
	}

	authorType := struct {
		TagName string `json:"tag_name"`
	}{}
	if err = json.Unmarshal(body, &authorType); err != nil {
		l.Println(fmt.Sprintf("failed to parse: %s", string(body)))
		return
	}

	currentV, err := version.NewVersion(config.Version)
	if err != nil {
		l.Println(err, "failed to parse current version")
		return
	}
	latestV, err := version.NewVersion(authorType.TagName)
	if err != nil {
		l.Println(err, "failed to parse latest version")
		return
	}

	if currentV.LessThan(latestV) {
		l.Printf("new version is available: %s, consider updating the client", coloredNotice(latestV))
	}
}

// getVersionRequest send a job request to service
func getVersionRequest(clientVer string, host string) (ver string, err error) {
	dialTimeoutCtx, dialCancel := context.WithTimeout(context.Background(), OptimusDialTimeout)
	defer dialCancel()

	var conn *grpc.ClientConn
	if conn, err = createConnection(dialTimeoutCtx, host); err != nil {
		return "", err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), versionTimeout)
	defer cancel()

	runtime := pb.NewRuntimeServiceClient(conn)

	versionResponse, err := runtime.Version(ctx, &pb.VersionRequest{
		Client: clientVer,
	})
	if err != nil {
		return "", errors.Wrapf(err, "request failed for version")
	}
	return versionResponse.Server, nil
}
