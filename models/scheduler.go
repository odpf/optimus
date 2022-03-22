package models

import (
	"context"
	"fmt"
	"time"

	"github.com/odpf/optimus/core/cron"

	"github.com/odpf/optimus/core/progress"
)

var (
	// BatchScheduler is a single unit initialized at the start of application
	// based on configs. This will be used to perform schedule triggered
	// operations to support target scheduling engine
	BatchScheduler SchedulerUnit

	// ManualScheduler is a single unit initialized at the start of application
	// based on configs. This will be used to execute one shot manual triggered
	// operations to support target scheduling engine
	ManualScheduler SchedulerUnit
)

// SchedulerUnit is implemented by supported schedulers
type SchedulerUnit interface {
	GetName() string

	VerifyJob(ctx context.Context, namespace NamespaceSpec, job JobSpec) error
	ListJobs(ctx context.Context, namespace NamespaceSpec, opts SchedulerListOptions) ([]Job, error)
	DeployJobs(ctx context.Context, namespace NamespaceSpec, jobs []JobSpec, obs progress.Observer) error
	DeleteJobs(ctx context.Context, namespace NamespaceSpec, jobNames []string, obs progress.Observer) error

	// Bootstrap will be executed per project when the application boots up
	// this can be used to do adhoc commands for initialization of scheduler
	Bootstrap(context.Context, ProjectSpec) error

	// GetJobStatus should return the current and previous status of job
	GetJobStatus(ctx context.Context, projSpec ProjectSpec, jobName string) ([]JobStatus, error)

	// Clear clears state of job between provided start and end dates
	Clear(ctx context.Context, projSpec ProjectSpec, jobName string, startDate, endDate time.Time) error

	// GetJobRunStatus should return batch of runs of a job
	GetJobRunStatus(ctx context.Context, projectSpec ProjectSpec, jobName string, startDate time.Time,
		endDate time.Time, batchSize int) ([]JobStatus, error)

	//GetJobRuns return all the job runs based on query
	GetJobRuns(ctx context.Context, projectSpec ProjectSpec, param *JobQuery, spec *cron.ScheduleSpec) ([]JobRun, error)
}

type SchedulerListOptions struct {
	OnlyName bool
}

type JobStatus struct {
	ScheduledAt time.Time
	State       JobRunState
}

// progress events
type (
	// EventJobSpecCompile represents a specification
	// being compiled to a Job
	EventJobSpecCompiled struct{ Name string }

	// EventJobUpload represents the compiled Job
	// being uploaded
	EventJobUpload struct {
		Name string
		Err  error
	}

	// EventJobRemoteDelete signifies that a
	// compiled job from a remote repository is being deleted
	EventJobRemoteDelete struct{ Name string }
)

func (e *EventJobSpecCompiled) String() string {
	return fmt.Sprintf("compiling: %s", e.Name)
}

func (e *EventJobUpload) String() string {
	if e.Err != nil {
		return fmt.Sprintf("uploading: %s, failed with error): %s", e.Name, e.Err.Error())
	}
	return fmt.Sprintf("uploaded: %s", e.Name)
}

func (e *EventJobRemoteDelete) String() string {
	return fmt.Sprintf("deleting: %s", e.Name)
}

// ExecutorUnit executes the actual job instance
type ExecutorUnit interface {
	// Start initiates the instance execution
	Start(ctx context.Context, req ExecutorStartRequest) (*ExecutorStartResponse, error)

	// Stop aborts the execution
	Stop(ctx context.Context, req ExecutorStopRequest) error

	// WaitForFinish returns a channel that should return the exit code of execution
	// once it finishes
	WaitForFinish(ctx context.Context, id string) (chan int, error)

	// Stats provides current statistics of the running/finished instance
	Stats(ctx context.Context, id string) (*ExecutorStats, error)
}

type ExecutorStartRequest struct {
	// ID will be used for identifying the job in future calls
	ID string

	Job       JobSpec
	Namespace NamespaceSpec
}

type ExecutorStopRequest struct {
	ID     string
	Signal string
}

type ExecutorStartResponse struct {
}

type ExecutorStats struct {
	Logs   []byte
	Status string
}
