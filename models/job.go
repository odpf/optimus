package models

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/odpf/optimus/core/progress"
)

var (
	ErrNoSuchSpec  = errors.New("spec not found")
	ErrNoDAGSpecs  = errors.New("no job specifications found")
	ErrNoSuchJob   = errors.New("job not found")
	ErrNoJobs      = errors.New("no job found")
	ErrNoResources = errors.New("no resources found")
	ErrNoSuchAsset = errors.New("asset not found")
	ErrNoSuchHook  = errors.New("hook not found")
)

const (
	JobDatetimeLayout = "2006-01-02"

	// assuming all month are 30 days long for simplicity
	HoursInMonth = time.Duration(30) * 24 * time.Hour
	HoursInDay   = 24 * time.Hour

	// within a project
	JobSpecDependencyTypeIntra JobSpecDependencyType = "intra"
	// within optimus but cross project
	JobSpecDependencyTypeInter JobSpecDependencyType = "inter"
	// outside optimus
	JobSpecDependencyTypeExtra JobSpecDependencyType = "extra"

	JobEventTypeSLAMiss JobEventType = "sla_miss"
	JobEventTypeFailure JobEventType = "failure"
)

// JobSpec represents a job
// internal representation of the job
type JobSpec struct {
	ID                   uuid.UUID
	Version              int
	Name                 string
	Description          string
	Labels               map[string]string
	Owner                string
	Schedule             JobSpecSchedule
	Behavior             JobSpecBehavior
	Task                 JobSpecTask
	Dependencies         map[string]JobSpecDependency // job name to dependency
	Assets               JobAssets
	Hooks                []JobSpecHook
	Metadata             JobSpecMetadata
	ExternalDependencies ExternalDependency // external dependencies for http
}

func (js JobSpec) GetName() string {
	return js.Name
}

func (js JobSpec) GetHookByName(name string) (JobSpecHook, error) {
	for _, hook := range js.Hooks {
		if hook.Unit.Info().Name == name {
			return hook, nil
		}
	}
	return JobSpecHook{}, ErrNoSuchHook
}

func (js JobSpec) GetLabelsAsString() string {
	labels := ""
	for k, v := range js.Labels {
		labels += fmt.Sprintf("%s=%s,", strings.TrimSpace(k), strings.TrimSpace(v))
	}
	return strings.TrimRight(labels, ",")
}

type JobSpecSchedule struct {
	StartDate time.Time
	EndDate   *time.Time
	Interval  string // could be empty string for no schedule
}

type JobSpecBehavior struct {
	DependsOnPast bool
	CatchUp       bool
	Retry         JobSpecBehaviorRetry
	Notify        []JobSpecNotifier
}

type JobSpecBehaviorRetry struct {
	Count              int
	Delay              time.Duration
	ExponentialBackoff bool
}

type JobSpecNotifier struct {
	On       JobEventType
	Config   map[string]string
	Channels []string
}

type JobSpecTask struct {
	Unit     *Plugin `json:"-" yaml:"-"`
	Config   JobSpecConfigs
	Window   JobSpecTaskWindow
	Priority int
}

type JobSpecTaskDestination struct {
	Destination string
	Type        DestinationType
}

type JobSpecTaskDependencies []string

// using array to keep order, map would be more performant
type JobSpecConfigs []JobSpecConfigItem

func (j JobSpecConfigs) Get(name string) (string, bool) {
	for _, conf := range j {
		if conf.Name == name {
			return conf.Value, true
		}
	}
	return "", false
}

type JobSpecConfigItem struct {
	Name  string
	Value string
}

type JobSpecTaskWindow struct {
	Size       time.Duration
	Offset     time.Duration
	TruncateTo string
}

func (w *JobSpecTaskWindow) GetStart(scheduledAt time.Time) time.Time {
	s, _ := w.getWindowDate(scheduledAt, w.Size, w.Offset, w.TruncateTo)
	return s
}

func (w *JobSpecTaskWindow) GetEnd(scheduledAt time.Time) time.Time {
	_, e := w.getWindowDate(scheduledAt, w.Size, w.Offset, w.TruncateTo)
	return e
}

func (w *JobSpecTaskWindow) getWindowDate(today time.Time, windowSize, windowOffset time.Duration, windowTruncateTo string) (time.Time, time.Time) {
	floatingEnd := today

	// apply truncation to end
	if windowTruncateTo == "h" {
		// remove time upto hours
		floatingEnd = floatingEnd.Truncate(time.Hour)
	} else if windowTruncateTo == "d" {
		// remove time upto day
		floatingEnd = floatingEnd.Truncate(HoursInDay)
	} else if windowTruncateTo == "w" {
		// shift current window to nearest Sunday
		nearestSunday := time.Duration(time.Saturday-floatingEnd.Weekday()+1) * HoursInDay
		floatingEnd = floatingEnd.Add(nearestSunday)
		floatingEnd = floatingEnd.Truncate(HoursInDay)
	}

	windowEnd := floatingEnd.Add(windowOffset)
	windowStart := windowEnd.Add(-windowSize)

	// handle monthly windows separately as every month is not of same size
	if windowTruncateTo == "M" {
		floatingEnd = today
		// shift current window to nearest month start and end

		// truncate the date
		floatingEnd = time.Date(floatingEnd.Year(), floatingEnd.Month(), 1, 0, 0, 0, 0, time.UTC)

		// then add the month offset
		// for handling offset, treat 30 days as 1 month
		offsetMonths := windowOffset / HoursInMonth
		floatingEnd = floatingEnd.AddDate(0, int(offsetMonths), 0)

		// then find the last day of this month
		floatingEnd = floatingEnd.AddDate(0, 1, -1)

		// final end is computed
		windowEnd = floatingEnd.Truncate(HoursInDay)

		// truncate days/hours from window start as well
		floatingStart := time.Date(floatingEnd.Year(), floatingEnd.Month(), 1, 0, 0, 0, 0, time.UTC)
		// for handling size, treat 30 days as 1 month, and as we have already truncated current month
		// subtract 1 from this
		sizeMonths := (windowSize / HoursInMonth) - 1
		if sizeMonths > 0 {
			floatingStart = floatingStart.AddDate(0, int(-sizeMonths), 0)
		}

		// final start is computed
		windowStart = floatingStart
	}

	return windowStart, windowEnd
}

type JobSpecHook struct {
	Config    JobSpecConfigs
	Unit      *Plugin
	DependsOn []*JobSpecHook
}

type JobSpecAsset struct {
	Name  string
	Value string
}

type JobAssets struct {
	data []JobSpecAsset
}

func (a JobAssets) FromMap(mp map[string]string) JobAssets {
	if len(mp) == 0 {
		return JobAssets{}
	}
	assets := JobAssets{
		data: make([]JobSpecAsset, 0),
	}
	for name, val := range mp {
		assets.data = append(assets.data, JobSpecAsset{
			Name:  name,
			Value: val,
		})
	}
	return assets
}

func (a *JobAssets) ToMap() map[string]string {
	if len(a.data) == 0 {
		return nil
	}
	mp := map[string]string{}
	for _, asset := range a.data {
		mp[asset.Name] = asset.Value
	}
	return mp
}

func (a *JobAssets) GetAll() []JobSpecAsset {
	return a.data
}

func (a JobAssets) New(data []JobSpecAsset) *JobAssets {
	return &JobAssets{
		data: data,
	}
}

func (a *JobAssets) GetByName(name string) (JobSpecAsset, error) {
	for _, asset := range a.data {
		if name == asset.Name {
			return asset, nil
		}
	}
	return JobSpecAsset{}, ErrNoSuchAsset
}

func (w *JobSpecTaskWindow) SizeString() string {
	return w.inHrs(int(w.Size.Hours()))
}

func (w *JobSpecTaskWindow) OffsetString() string {
	return w.inHrs(int(w.Offset.Hours()))
}

func (w *JobSpecTaskWindow) inHrs(hrs int) string {
	if hrs == 0 {
		return "0"
	}
	return fmt.Sprintf("%dh", hrs)
}

func (w *JobSpecTaskWindow) String() string {
	return fmt.Sprintf("size_%dh", int(w.Size.Hours()))
}

type JobSpecDependencyType string

func (j JobSpecDependencyType) String() string {
	return string(j)
}

type JobSpecDependency struct {
	Project *ProjectSpec
	Job     *JobSpec
	Type    JobSpecDependencyType
}

type ExternalDependency struct {
	HTTPDependencies []HTTPDependency
}

type HTTPDependency struct {
	Name          string
	RequestParams map[string]string
	URL           string
	Headers       map[string]string
}

// JobService provides a high-level operations on DAGs
type JobService interface {
	// Create constructs a Job and commits it to a storage
	Create(context.Context, NamespaceSpec, JobSpec) error
	// GetByName fetches a Job by name for a specific namespace
	GetByName(context.Context, string, NamespaceSpec) (JobSpec, error)
	// KeepOnly deletes all jobs except the ones provided for a namespace
	KeepOnly(context.Context, NamespaceSpec, []JobSpec, progress.Observer) error
	// GetAll reads all job specifications of the given namespace
	GetAll(context.Context, NamespaceSpec) ([]JobSpec, error)
	// Delete deletes a job spec from all repos
	Delete(context.Context, NamespaceSpec, JobSpec) error
	// GetTaskDependencies returns job task dependency mod details
	GetTaskDependencies(context.Context, NamespaceSpec, JobSpec) (JobSpecTaskDestination,
		JobSpecTaskDependencies, error)

	// Run creates a new job run for provided job spec and schedules it to execute
	// immediately
	Run(context.Context, NamespaceSpec, []JobSpec, progress.Observer) error

	// GetByNameForProject fetches a Job by name for a specific project
	GetByNameForProject(context.Context, string, ProjectSpec) (JobSpec, NamespaceSpec, error)
	Sync(context.Context, NamespaceSpec, progress.Observer) error
	Check(context.Context, NamespaceSpec, []JobSpec, progress.Observer) error
	// ReplayDryRun returns the execution tree of jobSpec and its dependencies between start and endDate, and the ignored jobs
	ReplayDryRun(context.Context, ReplayRequest) (ReplayPlan, error)
	// Replay replays the jobSpec and its dependencies between start and endDate
	Replay(context.Context, ReplayRequest) (ReplayResult, error)
	// GetReplayStatus of a replay using its ID
	GetReplayStatus(context.Context, ReplayRequest) (ReplayState, error)
	// GetReplayList of a project
	GetReplayList(ctx context.Context, projectID ProjectID) ([]ReplaySpec, error)
	// GetByDestination fetches a Job by destination for a specific project
	GetByDestination(ctx context.Context, projectSpec ProjectSpec, destination string) (JobSpec, error)
	// GetDownstream fetches downstream jobspecs
	GetDownstream(ctx context.Context, projectSpec ProjectSpec, jobName string) ([]JobSpec, error)
	// Refresh Redeploy current persisted state of jobs
	Refresh(context.Context, ProjectSpec, []NamespaceJobNamePair, progress.Observer) error
}

// JobCompiler takes template file of a scheduler and after applying
// variables generates a executable input for scheduler.
type JobCompiler interface {
	Compile([]byte, NamespaceSpec, JobSpec) (Job, error)
}

// Job represents a compiled consumable item for scheduler
// this is generated from JobSpec
type Job struct {
	Name     string
	Contents []byte
}

type JobEventType string

// JobEvent refers to status updates related to job
// posted by scheduler
type JobEvent struct {
	Type  JobEventType
	Value map[string]*structpb.Value
}

type NotifyAttrs struct {
	Namespace NamespaceSpec

	JobSpec  JobSpec
	JobEvent JobEvent

	Route string
}

type Notifier interface {
	io.Closer
	Notify(ctx context.Context, attr NotifyAttrs) error
}

// JobSpecMetadata contains metadata for a job spec
type JobSpecMetadata struct {
	Resource JobSpecResource
}

// JobSpecResource represents resource management configuration
type JobSpecResource struct {
	Request JobSpecResourceConfig
	Limit   JobSpecResourceConfig
}

// JobSpecResourceConfig represents the resource configuration
type JobSpecResourceConfig struct {
	Memory string
	CPU    string
}

// JobQuery  represents the query to get run status from scheduler
type JobQuery struct {
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Filter      []string
	OnlyLastRun bool
}

type JobIDDependenciesPair struct {
	JobID            uuid.UUID
	DependentProject ProjectSpec
	DependentJobID   uuid.UUID
	Type             JobSpecDependencyType
}
