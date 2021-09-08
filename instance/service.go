package instance

import (
	"context"
	"time"

	"github.com/odpf/optimus/models"
	"github.com/odpf/optimus/store"
	"github.com/pkg/errors"
)

const (
	// these configs can be used as macros in task/hook config and job assets
	ConfigKeyDstart        = "DSTART"
	ConfigKeyDend          = "DEND"
	ConfigKeyExecutionTime = "EXECUTION_TIME"
	ConfigKeyDestination   = "JOB_DESTINATION"
)

type InstanceSpecRepoFactory interface {
	New(models.JobSpec) store.InstanceSpecRepository
}

type Service struct {
	repoFac        InstanceSpecRepoFactory
	Now            func() time.Time
	templateEngine models.TemplateEngine
}

func (s *Service) Compile(namespace models.NamespaceSpec, jobSpec models.JobSpec, instanceSpec models.InstanceSpec,
	runType models.InstanceType, runName string) (envMap map[string]string, fileMap map[string]string, err error) {
	return NewContextManager(
		namespace, jobSpec, s.templateEngine).Generate(
		instanceSpec, runType, runName,
	)
}

func (s *Service) Register(jobSpec models.JobSpec, scheduledAt time.Time,
	instanceType models.InstanceType) (models.InstanceSpec, error) {
	jobRunRepo := s.repoFac.New(jobSpec)
	instanceToSave, err := s.PrepInstance(jobSpec, scheduledAt)
	if err != nil {
		return models.InstanceSpec{}, errors.Wrap(err, "failed to register instance")
	}

	switch instanceType {
	case models.InstanceTypeTask:
		// clear and save fresh
		if err := jobRunRepo.Clear(scheduledAt); err != nil && !errors.Is(err, store.ErrResourceNotFound) {
			return models.InstanceSpec{}, errors.Wrapf(err, "failed to clear instance of job %s",
				scheduledAt.String())
		}
		if err := jobRunRepo.Save(instanceToSave); err != nil {
			return models.InstanceSpec{}, err
		}
	case models.InstanceTypeHook:
		// store only if not already exists
		_, err := jobRunRepo.GetByScheduledAt(scheduledAt)
		if errors.Is(err, store.ErrResourceNotFound) {
			if err := jobRunRepo.Save(instanceToSave); err != nil {
				return models.InstanceSpec{}, err
			}
		} else if err != nil {
			return models.InstanceSpec{}, err
		}

	default:
		return models.InstanceSpec{}, errors.Errorf("invalid instance type: %s", instanceType)
	}

	// get whatever is saved, querying again ensures it was saved correctly
	instanceSpec, err := jobRunRepo.GetByScheduledAt(scheduledAt)
	if err != nil {
		return models.InstanceSpec{}, errors.Wrapf(err, "failed to save instance scheduled at: %s", scheduledAt.String())
	}
	return instanceSpec, nil
}

func (s *Service) PrepInstance(jobSpec models.JobSpec, scheduledAt time.Time) (models.InstanceSpec, error) {
	var jobDestination string
	if jobSpec.Task.Unit.DependencyMod != nil {
		jobDestinationResponse, err := jobSpec.Task.Unit.DependencyMod.GenerateDestination(context.TODO(), models.GenerateDestinationRequest{
			Config: models.PluginConfigs{}.FromJobSpec(jobSpec.Task.Config),
			Assets: models.PluginAssets{}.FromJobSpec(jobSpec.Assets),
		})
		if err != nil {
			return models.InstanceSpec{}, errors.Wrapf(err, "failed to generate destination for job %s", jobSpec.Name)
		}
		jobDestination = jobDestinationResponse.Destination
	}

	return models.InstanceSpec{
		Job:         jobSpec,
		ScheduledAt: scheduledAt,
		State:       models.InstanceStateRunning,

		// append optimus configs based on the values of a specific JobRun eg, jobScheduledTime
		Data: []models.InstanceSpecData{
			{
				Name:  ConfigKeyExecutionTime,
				Value: s.Now().Format(models.InstanceScheduledAtTimeLayout),
				Type:  models.InstanceDataTypeEnv,
			},
			{
				Name:  ConfigKeyDstart,
				Value: jobSpec.Task.Window.GetStart(scheduledAt).Format(models.InstanceScheduledAtTimeLayout),
				Type:  models.InstanceDataTypeEnv,
			},
			{
				Name:  ConfigKeyDend,
				Value: jobSpec.Task.Window.GetEnd(scheduledAt).Format(models.InstanceScheduledAtTimeLayout),
				Type:  models.InstanceDataTypeEnv,
			},
			{
				Name:  ConfigKeyDestination,
				Value: jobDestination,
				Type:  models.InstanceDataTypeEnv,
			},
		},
	}, nil
}

func NewService(repoFac InstanceSpecRepoFactory, timeFunc func() time.Time, te models.TemplateEngine) *Service {
	return &Service{
		repoFac:        repoFac,
		Now:            timeFunc,
		templateEngine: te,
	}
}
