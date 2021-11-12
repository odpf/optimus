package meta

import (
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	pb "github.com/odpf/optimus/api/proto/odpf/metadata/optimus/v1"
	"github.com/odpf/optimus/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobAdapter struct {
}

func (a JobAdapter) buildUrn(projectSpec models.ProjectSpec, jobSpec models.JobSpec) string {
	return fmt.Sprintf("%s::job/%s", projectSpec.Name, jobSpec.Name)
}

func (a JobAdapter) FromJobSpec(namespaceSpec models.NamespaceSpec, jobSpec models.JobSpec) (*models.JobMetadata, error) {
	taskSchema := jobSpec.Task.Unit.Info()

	taskPluginConfigs := models.PluginConfigs{}
	for _, c := range jobSpec.Task.Config {
		taskPluginConfigs = append(taskPluginConfigs, models.PluginConfig{
			Name:  c.Name,
			Value: c.Value,
		})
	}
	taskPluginAssets := models.PluginAssets{}
	for _, c := range jobSpec.Assets.GetAll() {
		taskPluginAssets = append(taskPluginAssets, models.PluginAsset{
			Name:  c.Name,
			Value: c.Value,
		})
	}
	var jobDestination string
	if jobSpec.Task.Unit.DependencyMod != nil {
		jobDestinationResponse, err := jobSpec.Task.Unit.DependencyMod.GenerateDestination(context.TODO(), models.GenerateDestinationRequest{
			Config: models.PluginConfigs{}.FromJobSpec(jobSpec.Task.Config),
			Assets: models.PluginAssets{}.FromJobSpec(jobSpec.Assets),
		})
		if err != nil {
			return nil, err
		}
		jobDestination = jobDestinationResponse.URN()
	}

	taskMetadata := models.JobTaskMetadata{
		Name:        taskSchema.Name,
		Image:       taskSchema.Image,
		Description: taskSchema.Description,
		Destination: jobDestination,
		Config:      jobSpec.Task.Config,
		Window:      jobSpec.Task.Window,
		Priority:    jobSpec.Task.Priority,
	}

	resourceMetadata := models.JobMetadata{
		Urn:          a.buildUrn(namespaceSpec.ProjectSpec, jobSpec),
		Name:         jobSpec.Name,
		Tenant:       namespaceSpec.ProjectSpec.Name,
		Namespace:    namespaceSpec.Name,
		Version:      jobSpec.Version,
		Description:  jobSpec.Description,
		Labels:       CompileSpecLabels(jobSpec),
		Owner:        jobSpec.Owner,
		Task:         taskMetadata,
		Schedule:     jobSpec.Schedule,
		Behavior:     jobSpec.Behavior,
		Dependencies: []models.JobDependencyMetadata{},
		Hooks:        []models.JobHookMetadata{},
	}

	for _, depJob := range jobSpec.Dependencies {
		resourceMetadata.Dependencies = append(resourceMetadata.Dependencies, models.JobDependencyMetadata{
			Tenant: depJob.Project.Name,
			Job:    depJob.Job.Name,
			Type:   depJob.Type.String(),
		})
	}

	for _, hook := range jobSpec.Hooks {
		schema := hook.Unit.Info()
		resourceMetadata.Hooks = append(resourceMetadata.Hooks, models.JobHookMetadata{
			Name:        schema.Name,
			Image:       schema.Image,
			Description: schema.Description,
			Config:      hook.Config,
			Type:        schema.HookType,
			DependsOn:   schema.DependsOn,
		})
	}

	return &resourceMetadata, nil
}

func (a JobAdapter) CompileKey(urn string) ([]byte, error) {
	return proto.Marshal(&pb.JobMetadataKey{
		Urn: urn,
	})
}

func (a JobAdapter) CompileMessage(jobMetadata *models.JobMetadata) ([]byte, error) {
	timestamp := timestamppb.New(time.Now())

	jobSchedule, err := a.compileJobSchedule(jobMetadata)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(&pb.JobMetadata{
		Urn:         jobMetadata.Urn,
		Name:        jobMetadata.Name,
		Tenant:      jobMetadata.Tenant,
		Namespace:   jobMetadata.Namespace,
		Version:     int32(jobMetadata.Version),
		Description: jobMetadata.Description,
		Labels:      a.compileProtoLabels(jobMetadata),
		Owner:       jobMetadata.Owner,
		Task:        a.compileTask(jobMetadata),
		Schedule:    jobSchedule,
		Behaviour: &pb.JobBehavior{
			DependsOnPast: jobMetadata.Behavior.DependsOnPast,
			Catchup:       jobMetadata.Behavior.CatchUp,
		},
		Hooks:          a.compileHooks(jobMetadata),
		Dependencies:   a.compileDependency(jobMetadata),
		EventTimestamp: timestamp,
	})
}

func (a JobAdapter) compileTask(resource *models.JobMetadata) *pb.JobTask {
	var taskConfig []*pb.JobTaskConfig
	for _, config := range resource.Task.Config {
		taskConfig = append(taskConfig, &pb.JobTaskConfig{
			Name:  config.Name,
			Value: config.Value,
		})
	}

	taskWindow := &pb.JobTaskWindow{
		Size:       resource.Task.Window.Size.String(),
		Offset:     resource.Task.Window.Offset.String(),
		TruncateTo: resource.Task.Window.TruncateTo,
	}

	return &pb.JobTask{
		Name:        resource.Task.Name,
		Image:       resource.Task.Image,
		Description: resource.Task.Description,
		Destination: resource.Task.Destination,
		Config:      taskConfig,
		Window:      taskWindow,
		Priority:    int32(resource.Task.Priority),
	}
}

func (a JobAdapter) compileHooks(resource *models.JobMetadata) (hooks []*pb.JobHook) {
	for _, hook := range resource.Hooks {
		var hookConfig []*pb.JobHookConfig
		for _, config := range hook.Config {
			hookConfig = append(hookConfig, &pb.JobHookConfig{
				Name:  config.Name,
				Value: config.Value,
			})
		}
		hooks = append(hooks, &pb.JobHook{
			Name:        hook.Name,
			Image:       hook.Image,
			Description: hook.Description,
			Config:      hookConfig,
			Type:        hook.Type.String(),
			DependsOn:   hook.DependsOn,
		})
	}
	return
}

func (a JobAdapter) compileJobSchedule(resource *models.JobMetadata) (*pb.JobSchedule, error) {
	scheduleStartDate := timestamppb.New(resource.Schedule.StartDate)

	var scheduleEndDate *timestamppb.Timestamp
	if resource.Schedule.EndDate != nil {
		scheduleEndDate = timestamppb.New(*resource.Schedule.EndDate)
	}

	return &pb.JobSchedule{
		StartDate: scheduleStartDate,
		EndDate:   scheduleEndDate,
		Interval:  resource.Schedule.Interval,
	}, nil
}

func (a JobAdapter) compileDependency(resource *models.JobMetadata) (dependencies []*pb.JobDependency) {
	for _, dependency := range resource.Dependencies {
		dependencies = append(dependencies, &pb.JobDependency{
			Tenant: dependency.Tenant,
			Job:    dependency.Job,
			Type:   dependency.Type,
		})
	}
	return
}

func (a JobAdapter) compileProtoLabels(resource *models.JobMetadata) (labels []*pb.JobLabel) {
	for _, config := range resource.Labels {
		labels = append(labels, &pb.JobLabel{
			Name:  config.Name,
			Value: config.Value,
		})
	}
	return
}

func CompileSpecLabels(resource models.JobSpec) (labels []models.JobMetadataLabelItem) {
	for k, v := range resource.Labels {
		labels = append(labels, models.JobMetadataLabelItem{
			Name:  k,
			Value: v,
		})
	}
	return
}
