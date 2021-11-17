package postgres

import (
	"context"
	"encoding/json"
	"time"

	"github.com/odpf/optimus/store"

	"github.com/google/uuid"
	"github.com/odpf/optimus/models"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Namespace struct {
	ID     uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name   string    `gorm:"not null;unique"`
	Config datatypes.JSON

	ProjectID uuid.UUID
	Project   Project `gorm:"foreignKey:ProjectID"`

	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func (p Namespace) FromSpec(spec models.NamespaceSpec) (Namespace, error) {
	jsonBytes, err := json.Marshal(spec.Config)
	if err != nil {
		return Namespace{}, nil
	}

	return Namespace{
		ID:     spec.ID,
		Name:   spec.Name,
		Config: jsonBytes,
	}, nil
}

func (p Namespace) FromSpecWithProject(spec models.NamespaceSpec, proj models.ProjectSpec) (Namespace, error) {
	adaptNamespace, err := p.FromSpec(spec)
	if err != nil {
		return adaptNamespace, err
	}

	adaptProject, err := Project{}.FromSpec(proj)
	if err != nil {
		return adaptNamespace, err
	}

	adaptNamespace.Project = adaptProject
	adaptNamespace.ProjectID = adaptProject.ID

	return adaptNamespace, nil
}

func (p Namespace) ToSpec(project models.ProjectSpec) (models.NamespaceSpec, error) {
	var conf map[string]string
	if err := json.Unmarshal(p.Config, &conf); err != nil {
		return models.NamespaceSpec{}, err
	}

	return models.NamespaceSpec{
		ID:          p.ID,
		Name:        p.Name,
		Config:      conf,
		ProjectSpec: project,
	}, nil
}

func (p Namespace) ToSpecWithProjectSecrets(hash models.ApplicationKey) (models.NamespaceSpec, error) {
	var conf map[string]string
	if err := json.Unmarshal(p.Config, &conf); err != nil {
		return models.NamespaceSpec{}, err
	}

	pSpec, err := p.Project.ToSpecWithSecrets(hash)
	if err != nil {
		return models.NamespaceSpec{}, err
	}
	return models.NamespaceSpec{
		ID:          p.ID,
		Name:        p.Name,
		Config:      conf,
		ProjectSpec: pSpec,
	}, nil
}

type namespaceRepository struct {
	db      *gorm.DB
	project models.ProjectSpec
	hash    models.ApplicationKey
}

func (repo *namespaceRepository) Insert(ctx context.Context, resource models.NamespaceSpec) error {
	c, err := Namespace{}.FromSpecWithProject(resource, repo.project)
	if err != nil {
		return err
	}
	if len(c.Name) == 0 {
		return errors.New("name cannot be empty")
	}
	return repo.db.WithContext(ctx).Create(&c).Error
}

func (repo *namespaceRepository) Save(ctx context.Context, spec models.NamespaceSpec) error {
	existingResource, err := repo.GetByName(ctx, spec.Name)
	if errors.Is(err, store.ErrResourceNotFound) {
		return repo.Insert(ctx, spec)
	} else if err != nil {
		return errors.Wrap(err, "unable to find namespace by name")
	}
	if len(spec.Config) == 0 {
		return store.ErrEmptyConfig
	}
	resource, err := Namespace{}.FromSpec(spec)
	if err != nil {
		return err
	}
	resource.ID = existingResource.ID
	return repo.db.WithContext(ctx).Model(resource).Updates(resource).Error
}

func (repo *namespaceRepository) GetByName(ctx context.Context, name string) (models.NamespaceSpec, error) {
	var r Namespace
	if err := repo.db.WithContext(ctx).Preload("Project").Preload("Project.Secrets").Where("name = ? AND project_id = ?", name, repo.project.ID).First(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.NamespaceSpec{}, store.ErrResourceNotFound
		}
		return models.NamespaceSpec{}, err
	}
	return r.ToSpecWithProjectSecrets(repo.hash)
}

func (repo *namespaceRepository) GetAll(ctx context.Context) ([]models.NamespaceSpec, error) {
	var specs []models.NamespaceSpec
	var namespaces []Namespace
	if err := repo.db.WithContext(ctx).Preload("Project").Preload("Project.Secrets").Where("project_id = ?", repo.project.ID).Find(&namespaces).Error; err != nil {
		return specs, err
	}

	for _, namespace := range namespaces {
		adapt, err := namespace.ToSpecWithProjectSecrets(repo.hash)
		if err != nil {
			return specs, err
		}
		specs = append(specs, adapt)
	}
	return specs, nil
}

func NewNamespaceRepository(db *gorm.DB, project models.ProjectSpec, hash models.ApplicationKey) *namespaceRepository {
	return &namespaceRepository{
		db:      db,
		project: project,
		hash:    hash,
	}
}
