package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/datatypes"

	"github.com/odpf/optimus/store"

	"github.com/google/uuid"
	"github.com/odpf/optimus/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Resource struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`

	ProjectID uuid.UUID
	Project   Project `gorm:"foreignKey:ProjectID"`

	NamespaceID uuid.UUID
	Namespace   Namespace `gorm:"foreignKey:NamespaceID"`

	Version   int
	Name      string `gorm:"not null"`
	Type      string `gorm:"not null"`
	Datastore string `gorm:"not null"`
	URN       string `gorm:"not null"`

	Spec   []byte
	Assets datatypes.JSON
	Labels datatypes.JSON

	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func (r Resource) FromSpec(resourceSpec models.ResourceSpec) (Resource, error) {
	assetBytes, err := json.Marshal(resourceSpec.Assets)
	if err != nil {
		return Resource{}, err
	}
	labelBytes, err := json.Marshal(resourceSpec.Labels)
	if err != nil {
		return Resource{}, err
	}

	// serialize resource spec without assets to one of the datastore provided wire format
	controller, ok := resourceSpec.Datastore.Types()[resourceSpec.Type]
	if !ok {
		return Resource{}, fmt.Errorf("unknown type of datastore %s", resourceSpec.Type)
	}

	binaryReadySpec := resourceSpec
	binaryReadySpec.Assets = nil
	binaryReadySpec.Labels = nil
	serializedSpec, err := controller.Adapter().ToYaml(binaryReadySpec)
	if err != nil {
		return Resource{}, errors.Wrapf(err, "controller.Adapter().ToYaml: %v", binaryReadySpec)
	}

	urn, err := controller.GenerateURN(resourceSpec.Spec)
	if err != nil {
		return Resource{}, err
	}

	return Resource{
		ID:        resourceSpec.ID,
		Version:   resourceSpec.Version,
		Name:      resourceSpec.Name,
		Type:      resourceSpec.Type.String(),
		Datastore: resourceSpec.Datastore.Name(),
		URN:       urn,
		Spec:      serializedSpec,
		Assets:    assetBytes,
		Labels:    labelBytes,
	}, nil
}

func (r Resource) FromSpecWithNamespace(resourceSpec models.ResourceSpec, namespace models.NamespaceSpec) (Resource, error) {
	adaptResource, err := r.FromSpec(resourceSpec)
	if err != nil {
		return Resource{}, err
	}

	// namespace
	adaptNamespace, err := Namespace{}.FromSpecWithProject(namespace, namespace.ProjectSpec)
	if err != nil {
		return Resource{}, err
	}
	adaptResource.NamespaceID = adaptNamespace.ID
	adaptResource.Namespace = adaptNamespace

	// project
	adaptProject, err := Project{}.FromSpec(namespace.ProjectSpec)
	if err != nil {
		return Resource{}, err
	}
	adaptResource.ProjectID = adaptProject.ID
	adaptResource.Project = adaptProject

	return adaptResource, nil
}

func (r Resource) ToSpec(ds models.Datastorer) (models.ResourceSpec, error) {
	resourceType := models.ResourceType(r.Type)

	// deserialize resource spec without assets to one of the datastore provided wire format
	controller, ok := ds.Types()[resourceType]
	if !ok {
		return models.ResourceSpec{}, fmt.Errorf("unknown type of datastore %s", resourceType)
	}
	deserializedSpec, err := controller.Adapter().FromYaml(r.Spec)
	if err != nil {
		return models.ResourceSpec{}, errors.Wrapf(err, "controller.Adapter().FromYaml: %s", string(r.Spec))
	}

	var assets map[string]string
	if err := json.Unmarshal(r.Assets, &assets); err != nil {
		return models.ResourceSpec{}, err
	}
	var labels map[string]string
	if err := json.Unmarshal(r.Labels, &labels); err != nil {
		return models.ResourceSpec{}, err
	}

	return models.ResourceSpec{
		ID:        r.ID,
		Version:   r.Version,
		Name:      r.Name,
		Type:      resourceType,
		Datastore: ds,
		URN:       r.URN,
		Spec:      deserializedSpec.Spec,
		Assets:    assets,
		Labels:    labels,
	}, nil
}

type projectResourceSpecRepository struct {
	db        *gorm.DB
	project   models.ProjectSpec
	datastore models.Datastorer
}

func (repo *projectResourceSpecRepository) GetByName(ctx context.Context, name string) (models.ResourceSpec, models.NamespaceSpec, error) {
	var r Resource
	if err := repo.db.WithContext(ctx).Preload("Namespace").Where("project_id = ? AND datastore = ? AND name = ?",
		repo.project.ID, repo.datastore.Name(), name).First(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ResourceSpec{}, models.NamespaceSpec{}, store.ErrResourceNotFound
		}
		return models.ResourceSpec{}, models.NamespaceSpec{}, err
	}

	resourceSpec, err := r.ToSpec(repo.datastore)
	if err != nil {
		return models.ResourceSpec{}, models.NamespaceSpec{}, err
	}

	namespaceSpec, err := r.Namespace.ToSpec(repo.project)
	if err != nil {
		return models.ResourceSpec{}, models.NamespaceSpec{}, err
	}

	return resourceSpec, namespaceSpec, nil
}

func (repo *projectResourceSpecRepository) GetAll(ctx context.Context) ([]models.ResourceSpec, error) {
	specs := []models.ResourceSpec{}
	resources := []Resource{}
	if err := repo.db.WithContext(ctx).Where("project_id = ? AND datastore = ?", repo.project.ID, repo.datastore.Name()).Find(&resources).Error; err != nil {
		return specs, err
	}
	for _, r := range resources {
		adapted, err := r.ToSpec(repo.datastore)
		if err != nil {
			return specs, errors.Wrap(err, "failed to adapt resource")
		}
		specs = append(specs, adapted)
	}
	return specs, nil
}

func NewProjectResourceSpecRepository(db *gorm.DB, project models.ProjectSpec, ds models.Datastorer) *projectResourceSpecRepository {
	return &projectResourceSpecRepository{
		db:        db,
		project:   project,
		datastore: ds,
	}
}

type resourceSpecRepository struct {
	db                      *gorm.DB
	namespace               models.NamespaceSpec
	datastore               models.Datastorer
	projectResourceSpecRepo store.ProjectResourceSpecRepository
}

func (repo *resourceSpecRepository) Insert(ctx context.Context, resource models.ResourceSpec) error {
	if len(resource.Name) == 0 {
		return errors.New("name cannot be empty")
	}
	p, err := Resource{}.FromSpecWithNamespace(resource, repo.namespace)
	if err != nil {
		return err
	}
	// if soft deleted earlier
	if err := repo.HardDelete(ctx, resource.Name); err != nil {
		return err
	}
	return repo.db.WithContext(ctx).Create(&p).Error
}

func (repo *resourceSpecRepository) Save(ctx context.Context, spec models.ResourceSpec) error {
	existingResource, namespaceSpec, err := repo.projectResourceSpecRepo.GetByName(ctx, spec.Name)
	if errors.Is(err, store.ErrResourceNotFound) {
		return repo.Insert(ctx, spec)
	} else if err != nil {
		return errors.Wrap(err, "unable to find resource by name")
	}

	if namespaceSpec.ID != repo.namespace.ID {
		return errors.New(fmt.Sprintf("resource %s already exists for the project %s", spec.Name, repo.namespace.ProjectSpec.Name))
	}

	resource, err := Resource{}.FromSpec(spec)
	if err != nil {
		return err
	}
	resource.ID = existingResource.ID

	return repo.db.WithContext(ctx).Model(&resource).Updates(&resource).Error
}

func (repo *resourceSpecRepository) GetByName(ctx context.Context, name string) (models.ResourceSpec, error) {
	var r Resource
	if err := repo.db.WithContext(ctx).Where("namespace_id = ? AND datastore = ? AND name = ?",
		repo.namespace.ID, repo.datastore.Name(), name).First(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ResourceSpec{}, store.ErrResourceNotFound
		}
		return models.ResourceSpec{}, err
	}
	return r.ToSpec(repo.datastore)
}

func (repo *resourceSpecRepository) GetByID(ctx context.Context, id uuid.UUID) (models.ResourceSpec, error) {
	var r Resource
	if err := repo.db.WithContext(ctx).Where("namespace_id = ? AND id = ?",
		repo.namespace.ID, id).First(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ResourceSpec{}, store.ErrResourceNotFound
		}
		return models.ResourceSpec{}, err
	}
	return r.ToSpec(repo.datastore)
}

func (repo *resourceSpecRepository) GetByURN(ctx context.Context, urn string) (models.ResourceSpec, error) {
	var r Resource
	if err := repo.db.WithContext(ctx).Where("namespace_id = ? AND datastore = ? AND urn = ?",
		repo.namespace.ID, repo.datastore.Name(), urn).First(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ResourceSpec{}, store.ErrResourceNotFound
		}
		return models.ResourceSpec{}, err
	}
	return r.ToSpec(repo.datastore)
}

func (repo *resourceSpecRepository) GetAll(ctx context.Context) ([]models.ResourceSpec, error) {
	specs := []models.ResourceSpec{}
	resources := []Resource{}
	if err := repo.db.WithContext(ctx).Where("namespace_id = ? AND datastore = ?", repo.namespace.ID, repo.datastore.Name()).Find(&resources).Error; err != nil {
		return specs, err
	}
	for _, r := range resources {
		adapted, err := r.ToSpec(repo.datastore)
		if err != nil {
			return specs, errors.Wrap(err, "failed to adapt resource")
		}
		specs = append(specs, adapted)
	}
	return specs, nil
}

func (repo *resourceSpecRepository) Delete(ctx context.Context, name string) error {
	return repo.db.WithContext(ctx).Where("namespace_id = ? AND datastore = ? AND name = ? ", repo.namespace.ID, repo.datastore.Name(), name).Delete(&Resource{}).Error
}

func (repo *resourceSpecRepository) HardDelete(ctx context.Context, name string) error {
	return repo.db.WithContext(ctx).Unscoped().Where("namespace_id = ? AND datastore = ? AND name = ? ", repo.namespace.ID, repo.datastore.Name(), name).Delete(&Resource{}).Error
}

func NewResourceSpecRepository(db *gorm.DB, namespace models.NamespaceSpec, ds models.Datastorer, projectResourceSpecRepo store.ProjectResourceSpecRepository) *resourceSpecRepository {
	return &resourceSpecRepository{
		db:                      db,
		namespace:               namespace,
		datastore:               ds,
		projectResourceSpecRepo: projectResourceSpecRepo,
	}
}
