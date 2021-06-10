package mock

import (
	"context"
	"time"

	"github.com/odpf/optimus/core/multi_root_tree"

	"github.com/odpf/optimus/core/progress"
	"github.com/odpf/optimus/models"
	"github.com/odpf/optimus/store"
	"github.com/odpf/optimus/store/local"
	"github.com/stretchr/testify/mock"
)

// ProjectJobSpecRepoFactory to manage job specs at project level
type ProjectJobSpecRepoFactory struct {
	mock.Mock
}

func (repo *ProjectJobSpecRepoFactory) New(proj models.ProjectSpec) store.ProjectJobSpecRepository {
	return repo.Called(proj).Get(0).(store.ProjectJobSpecRepository)
}

// JobSpecRepoFactory to store raw specs
type ProjectJobSpecRepository struct {
	mock.Mock
}

func (repo *ProjectJobSpecRepository) GetByName(name string) (models.JobSpec, models.NamespaceSpec, error) {
	args := repo.Called(name)
	if args.Get(0) != nil {
		return args.Get(0).(models.JobSpec), args.Get(1).(models.NamespaceSpec), args.Error(2)
	}
	return models.JobSpec{}, models.NamespaceSpec{}, args.Error(1)
}

func (repo *ProjectJobSpecRepository) GetAll() ([]models.JobSpec, error) {
	args := repo.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.JobSpec), args.Error(1)
	}
	return []models.JobSpec{}, args.Error(1)
}

func (repo *ProjectJobSpecRepository) GetByDestination(dest string) (models.JobSpec, models.ProjectSpec, error) {
	args := repo.Called(dest)
	if args.Get(0) != nil {
		return args.Get(0).(models.JobSpec), args.Get(1).(models.ProjectSpec), args.Error(2)
	}
	return models.JobSpec{}, models.ProjectSpec{}, args.Error(2)
}

// JobSpecRepoFactory to store raw specs at namespace level
type JobSpecRepoFactory struct {
	mock.Mock
}

func (repo *JobSpecRepoFactory) New(namespace models.NamespaceSpec) store.JobSpecRepository {
	return repo.Called(namespace).Get(0).(store.JobSpecRepository)
}

// JobSpecRepoFactory to store raw specs
type JobSpecRepository struct {
	mock.Mock
}

func (repo *JobSpecRepository) Save(t models.JobSpec) error {
	return repo.Called(t).Error(0)
}

func (repo *JobSpecRepository) GetByName(name string) (models.JobSpec, error) {
	args := repo.Called(name)
	if args.Get(0) != nil {
		return args.Get(0).(models.JobSpec), args.Error(1)
	}
	return models.JobSpec{}, args.Error(1)
}

func (repo *JobSpecRepository) Delete(name string) error {
	return repo.Called(name).Error(0)
}

func (repo *JobSpecRepository) GetAll() ([]models.JobSpec, error) {
	args := repo.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.JobSpec), args.Error(1)
	}
	return []models.JobSpec{}, args.Error(1)
}

func (repo *JobSpecRepository) GetByDestination(dest string) (models.JobSpec, models.ProjectSpec, error) {
	args := repo.Called(dest)
	if args.Get(0) != nil {
		return args.Get(0).(models.JobSpec), args.Get(1).(models.ProjectSpec), args.Error(2)
	}
	return models.JobSpec{}, models.ProjectSpec{}, args.Error(2)
}

// JobRepoFactory to store compiled specs
type JobRepoFactory struct {
	mock.Mock
}

func (repo *JobRepoFactory) New(ctx context.Context, proj models.ProjectSpec) (store.JobRepository, error) {
	args := repo.Called(ctx, proj)
	return args.Get(0).(store.JobRepository), args.Error(1)
}

// JobRepository to store compiled specs

type JobRepository struct {
	mock.Mock
}

func (repo *JobRepository) Save(ctx context.Context, t models.Job) error {
	return repo.Called(ctx, t).Error(0)
}

func (repo *JobRepository) GetByName(ctx context.Context, name string) (models.Job, error) {
	args := repo.Called(ctx, name)
	return args.Get(0).(models.Job), args.Error(1)
}

func (repo *JobRepository) GetAll(ctx context.Context) ([]models.Job, error) {
	args := repo.Called(ctx)
	return args.Get(0).([]models.Job), args.Error(1)
}

func (repo *JobRepository) ListNames(ctx context.Context, namespace models.NamespaceSpec) ([]string, error) {
	args := repo.Called(ctx, namespace)
	return args.Get(0).([]string), args.Error(1)
}

func (repo *JobRepository) Delete(ctx context.Context, namespace models.NamespaceSpec, name string) error {
	args := repo.Called(ctx, namespace, name)
	return args.Error(0)
}

type JobConfigLocalFactory struct {
	mock.Mock
}

func (fac *JobConfigLocalFactory) New(inputs models.JobSpec) (local.Job, error) {
	args := fac.Called(inputs)
	return args.Get(0).(local.Job), args.Error(1)
}

type JobService struct {
	mock.Mock
}

func (srv *JobService) Create(spec2 models.NamespaceSpec, spec models.JobSpec) error {
	args := srv.Called(spec, spec2)
	return args.Error(0)
}

func (srv *JobService) GetByName(s string, spec models.NamespaceSpec) (models.JobSpec, error) {
	args := srv.Called(s, spec)
	return args.Get(0).(models.JobSpec), args.Error(1)
}

func (srv *JobService) Dump(spec2 models.NamespaceSpec, spec3 models.JobSpec) (models.Job, error) {
	args := srv.Called(spec2, spec3)
	return args.Get(0).(models.Job), args.Error(1)
}

func (srv *JobService) KeepOnly(spec models.NamespaceSpec, specs []models.JobSpec, observer progress.Observer) error {
	args := srv.Called(spec, specs)
	return args.Error(0)
}

func (srv *JobService) GetAll(spec models.NamespaceSpec) ([]models.JobSpec, error) {
	args := srv.Called(spec)
	return args.Get(0).([]models.JobSpec), args.Error(1)
}

func (srv *JobService) GetByNameForProject(s string, spec models.ProjectSpec) (models.JobSpec, models.NamespaceSpec, error) {
	args := srv.Called(s, spec)
	return args.Get(0).(models.JobSpec), args.Get(1).(models.NamespaceSpec), args.Error(2)
}

func (srv *JobService) Sync(ctx context.Context, spec models.NamespaceSpec, observer progress.Observer) error {
	args := srv.Called(ctx, spec, observer)
	return args.Error(0)
}

func (j *JobService) Check(namespaceSpec models.NamespaceSpec, specs []models.JobSpec, observer progress.Observer) error {
	args := j.Called(namespaceSpec, specs, observer)
	return args.Error(0)
}

func (j *JobService) Delete(ctx context.Context, c models.NamespaceSpec, job models.JobSpec) error {
	args := j.Called(ctx, c, job)
	return args.Error(0)
}

func (j *JobService) Replay(namespace models.NamespaceSpec, jobSpec models.JobSpec, dryRun bool, start time.Time, end time.Time) (*multi_root_tree.TreeNode, error) {
	args := j.Called(namespace, jobSpec, dryRun, start, end)
	return args.Get(0).(*multi_root_tree.TreeNode), args.Error(1)
}

type Compiler struct {
	mock.Mock
}

func (srv *Compiler) Compile(namespace models.NamespaceSpec, jobSpec models.JobSpec) (models.Job, error) {
	args := srv.Called(namespace, jobSpec)
	return args.Get(0).(models.Job), args.Error(1)
}

type DependencyResolver struct {
	mock.Mock
}

func (srv *DependencyResolver) Resolve(projectSpec models.ProjectSpec, projectJobSpecRepo store.ProjectJobSpecRepository,
	jobSpec models.JobSpec, obs progress.Observer) (models.JobSpec, error) {
	args := srv.Called(projectSpec, projectJobSpecRepo, jobSpec, obs)
	return args.Get(0).(models.JobSpec), args.Error(1)
}

type PriorityResolver struct {
	mock.Mock
}

func (srv *PriorityResolver) Resolve(jobSpecs []models.JobSpec) ([]models.JobSpec, error) {
	args := srv.Called(jobSpecs)
	return args.Get(0).([]models.JobSpec), args.Error(1)
}
