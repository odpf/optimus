package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	// generic backup configurations
	ConfigTTL              = "ttl"
	ConfigIgnoreDownstream = "ignore_downstream"

	BackupSpecKeyURN = "URN"
)

type BackupResourceRequest struct {
	Resource   ResourceSpec
	BackupSpec BackupRequest
	BackupTime time.Time
}

type BackupResourceResponse struct {
	ResultURN  string
	ResultSpec interface{}
}

type BackupRequest struct {
	ID                          uuid.UUID
	ResourceName                string
	Project                     ProjectSpec
	Namespace                   NamespaceSpec
	Datastore                   string
	Description                 string
	AllowedDownstreamNamespaces []string
	Config                      map[string]string
	DryRun                      bool
}

type BackupPlan struct {
	Resources        []string
	IgnoredResources []string
}

type BackupDetail struct {
	URN  string
	Spec interface{}
}

type BackupResponse struct {
	ResourceURN string
	Result      BackupDetail
}

type BackupSpec struct {
	ID          uuid.UUID
	Resource    ResourceSpec
	Result      map[string]interface{}
	Description string
	Config      map[string]string
	CreatedAt   time.Time
}

type BackupResult struct {
	Resources        []string
	IgnoredResources []string
}

type BackupService interface {
	BackupResourceDryRun(ctx context.Context, backupRequest BackupRequest, jobSpecs []JobSpec) (BackupPlan, error)
	BackupResource(ctx context.Context, backupRequest BackupRequest, jobSpecs []JobSpec) (BackupResult, error)
	ListResourceBackups(ctx context.Context, projectSpec ProjectSpec, datastoreName string) ([]BackupSpec, error)
	GetResourceBackup(ctx context.Context, projectSpec ProjectSpec, datastoreName string, id uuid.UUID) (BackupSpec, error)
}
