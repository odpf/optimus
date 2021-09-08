package job

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/odpf/optimus/core/tree"
	"github.com/odpf/optimus/models"
	"github.com/odpf/optimus/store"
	"github.com/odpf/optimus/utils"
	"github.com/odpf/salt/log"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

var (
	// ErrRequestQueueFull signifies that the deployment manager's
	// request queue is full
	ErrRequestQueueFull = errors.New("request queue is full")
	// ErrConflictedJobRun signifies other replay job / dependency run is active or instance already running
	ErrConflictedJobRun = errors.New("conflicted job run found")
	//ReplayRunTimeout signifies type of replay failure caused by timeout
	ReplayRunTimeout = "long running replay timeout"
	// TimestampLogFormat format of a timestamp will be used in logs
	TimestampLogFormat = "2006-01-02T15:04:05+00:00"
	// schedulerBatchSize number of run instances to be checked per request
	schedulerBatchSize = 100
)

const (
	syncInterval = "@every 5m"
)

type ReplayManagerConfig struct {
	NumWorkers    int
	WorkerTimeout time.Duration
	RunTimeout    time.Duration
}

type ReplayWorkerFactory interface {
	New() ReplayWorker
}

type ReplaySyncer interface {
	Sync(context.Context, time.Duration) error
}

type ReplayValidator interface {
	Validate(context.Context, store.ReplaySpecRepository, models.ReplayRequest, *tree.TreeNode) error
}

// Manager for replaying operation(s).
// Offers an asynchronous interface to pipeline, with a fixed size request queue
// Each replay request is handled by a replay worker and the number of parallel replay workers
// can be provided through configuration.
type Manager struct {
	// wait group to synchronise on workers
	wg sync.WaitGroup

	uuidProvider utils.UUIDProvider
	config       ReplayManagerConfig

	// request queue, used by workers
	requestQ chan models.ReplayRequest

	replayWorkerFactory ReplayWorkerFactory
	replaySpecRepoFac   ReplaySpecRepoFactory
	scheduler           models.SchedulerUnit
	replayValidator     ReplayValidator
	replaySyncer        ReplaySyncer
	syncerScheduler     *cron.Cron
	l                   log.Logger

	workerCapacity int32
}

// Replay a request asynchronously, returns a replay id that can
// can be used to query its status
func (m *Manager) Replay(ctx context.Context, reqInput models.ReplayRequest) (string, error) {
	replaySpecRepo := m.replaySpecRepoFac.New()

	replayTree, err := prepareReplayExecutionTree(reqInput)
	if err != nil {
		return "", err
	}
	if err := m.replayValidator.Validate(ctx, replaySpecRepo, reqInput, replayTree); err != nil {
		return "", err
	}
	if reqInput.ID, err = m.uuidProvider.NewUUID(); err != nil {
		return "", err
	}

	// save replay request and mark status as accepted
	replay := models.ReplaySpec{
		ID:            reqInput.ID,
		Job:           reqInput.Job,
		StartDate:     reqInput.Start,
		EndDate:       reqInput.End,
		Status:        models.ReplayStatusAccepted,
		ExecutionTree: replayTree,
	}

	// could get cancelled later if queue is full
	if err = replaySpecRepo.Insert(&replay); err != nil {
		return "", err
	}

	select {
	case m.requestQ <- reqInput:
		return reqInput.ID.String(), nil
	default:
		// all workers busy, mark the inserted request as cancelled
		_ = replaySpecRepo.UpdateStatus(reqInput.ID, models.ReplayStatusCancelled, models.ReplayMessage{
			Type:    models.ReplayStatusCancelled,
			Message: ErrRequestQueueFull.Error(),
		})
		return "", ErrRequestQueueFull
	}
}

// start a worker goroutine that runs the replay in background
func (m *Manager) spawnServiceWorker(worker ReplayWorker) {
	defer m.wg.Done()
	atomic.AddInt32(&m.workerCapacity, 1)
	for reqInput := range m.requestQ {
		atomic.AddInt32(&m.workerCapacity, -1)

		m.l.Info("worker picked up the request", "request id", reqInput.ID)
		ctx, cancelCtx := context.WithTimeout(context.Background(), m.config.WorkerTimeout)
		if err := worker.Process(ctx, reqInput); err != nil {
			m.l.Error("worker failed to process", "error", err)
		}
		cancelCtx()

		atomic.AddInt32(&m.workerCapacity, 1)
	}
}

// SchedulerSyncer to sync for latest replay status
func (m *Manager) SchedulerSyncer() {
	m.l.Debug("start synchronizing replays...")
	ctx, cancelCtx := context.WithTimeout(context.Background(), m.config.WorkerTimeout)
	defer cancelCtx()
	if err := m.replaySyncer.Sync(ctx, m.config.RunTimeout); err != nil {
		m.l.Error("syncer failed to process", "error", err)
	}
	m.l.Debug("replays synced")
}

// GetReplay using UUID
func (m *Manager) GetReplay(replayUUID uuid.UUID) (models.ReplaySpec, error) {
	return m.replaySpecRepoFac.New().GetByID(replayUUID)
}

// GetRunsStatus
func (m *Manager) GetRunStatus(ctx context.Context, projectSpec models.ProjectSpec, startDate time.Time,
	endDate time.Time, jobName string) ([]models.JobStatus, error) {
	batchEndDate := endDate.AddDate(0, 0, 1).Add(time.Second * -1)
	return m.scheduler.GetDagRunStatus(ctx, projectSpec, jobName, startDate, batchEndDate, schedulerBatchSize)
}

//Close stops consuming any new request
func (m *Manager) Close() error {
	if m.requestQ != nil {
		//stop accepting any more requests
		close(m.requestQ)
	}
	//wait for request worker to finish
	m.wg.Wait()

	if m.syncerScheduler != nil {
		//wait for syncer to finish
		<-m.syncerScheduler.Stop().Done()
	}
	return nil
}

func (m *Manager) Init() {
	if m.replaySyncer != nil {
		_, err := m.syncerScheduler.AddFunc(syncInterval, m.SchedulerSyncer)
		if err != nil {
			m.l.Fatal("Failed to sync scheduler", "error", err)
		}
		m.syncerScheduler.Start()
	}

	m.l.Info("starting replay workers", "count", m.config.NumWorkers)
	for i := 0; i < m.config.NumWorkers; i++ {
		m.wg.Add(1)
		worker := m.replayWorkerFactory.New()
		go m.spawnServiceWorker(worker)
	}

	// wait until all workers are ready
	for {
		if int(atomic.LoadInt32(&m.workerCapacity)) == m.config.NumWorkers {
			break
		}
		time.Sleep(time.Millisecond * 50)
	}
}

// NewManager constructs a new instance of Manager
func NewManager(l log.Logger, workerFact ReplayWorkerFactory, replaySpecRepoFac ReplaySpecRepoFactory, uuidProvider utils.UUIDProvider,
	config ReplayManagerConfig, scheduler models.SchedulerUnit, validator ReplayValidator, syncer ReplaySyncer) *Manager {
	mgr := &Manager{
		l:                   l,
		replayWorkerFactory: workerFact,
		config:              config,
		requestQ:            make(chan models.ReplayRequest),
		replaySpecRepoFac:   replaySpecRepoFac,
		uuidProvider:        uuidProvider,
		scheduler:           scheduler,
		replayValidator:     validator,
		replaySyncer:        syncer,
		workerCapacity:      0,
		syncerScheduler: cron.New(cron.WithChain(
			cron.SkipIfStillRunning(cron.DefaultLogger),
		)),
	}
	mgr.Init()
	return mgr
}
