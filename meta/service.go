package meta

import (
	"fmt"

	"github.com/odpf/optimus/core/progress"
	"github.com/odpf/optimus/models"
	"github.com/pkg/errors"
)

type MetaSvcFactory interface {
	New() models.MetadataService
}

type Service struct {
	writer     models.MetadataWriter
	jobAdapter models.JobMetadataAdapter
}

func NewService(writer models.MetadataWriter, builder models.JobMetadataAdapter) *Service {
	return &Service{
		writer:     writer,
		jobAdapter: builder,
	}
}

func (service Service) Publish(namespaceSpec models.NamespaceSpec, jobSpecs []models.JobSpec, po progress.Observer) error {
	for _, jobSpec := range jobSpecs {
		resource, err := service.jobAdapter.FromJobSpec(namespaceSpec, jobSpec)
		if err != nil {
			return err
		}

		protoKey, err := service.jobAdapter.CompileKey(resource.Urn)
		if err != nil {
			return errors.Wrapf(err, "failed to compile metadata proto key: %s", resource.Urn)
		}

		protoMsg, err := service.jobAdapter.CompileMessage(resource)
		if err != nil {
			return errors.Wrapf(err, "failed to compile metadata proto message: %s", resource.Urn)
		}

		if err = service.writer.Write(protoKey, protoMsg); err != nil {
			return errors.Wrapf(err, "failed to write metadata message: %s", resource.Urn)
		}
	}
	if po != nil {
		po.Notify(&EventPublish{SpecCount: len(jobSpecs)})
	}
	return nil
}

// EventPublish represents a specification being published to
// meta event stream
type EventPublish struct {
	SpecCount int
}

func (e *EventPublish) String() string {
	return fmt.Sprintf("published %d jobs metadata to event stream", e.SpecCount)
}
