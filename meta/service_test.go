package meta_test

import (
	"errors"
	"testing"

	"github.com/odpf/optimus/meta"
	"github.com/odpf/optimus/mock"
	"github.com/odpf/optimus/models"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	projectSpec := models.ProjectSpec{
		Name: "humara-projectSpec",
		Config: map[string]string{
			"bucket": "gs://some_folder",
		},
	}

	namespaceSpec := models.NamespaceSpec{
		Name: "humara-namespaceSpec",
		Config: map[string]string{
			"bucket": "gs://some_folder",
		},
		ProjectSpec: projectSpec,
	}

	jobSpecs := []models.JobSpec{
		{
			Name: "job-1",
			Task: models.JobSpecTask{
				Unit: nil,
				Config: models.JobSpecConfigs{
					{
						Name:  "do",
						Value: "this",
					},
				},
			},
			Assets: *models.JobAssets{}.New(
				[]models.JobSpecAsset{
					{
						Name:  "query.sql",
						Value: "select * from 1",
					},
				}),
		},
	}

	t.Run("should publish the job specs metadata", func(t *testing.T) {
		resource := &models.JobMetadata{Urn: jobSpecs[0].Name}
		protoKey := []byte("key")
		protoMsg := []byte("message")

		builder := new(mock.MetaBuilder)
		builder.On("FromJobSpec", namespaceSpec, jobSpecs[0]).Return(resource, nil)
		builder.On("CompileKey", jobSpecs[0].Name).Return(protoKey, nil)
		builder.On("CompileMessage", resource).Return(protoMsg, nil)
		defer builder.AssertExpectations(t)

		writer := new(mock.MetaWriter)
		writer.On("Write", protoKey, protoMsg).Return(nil)
		defer writer.AssertExpectations(t)

		service := meta.NewService(writer, builder)
		err := service.Publish(namespaceSpec, jobSpecs, nil)
		assert.Nil(t, err)
	})

	t.Run("should return error if writing to kafka fails", func(t *testing.T) {
		resource := &models.JobMetadata{Urn: jobSpecs[0].Name}
		protoKey := []byte("key")
		protoMsg := []byte("message")

		builder := new(mock.MetaBuilder)
		builder.On("FromJobSpec", namespaceSpec, jobSpecs[0]).Return(resource, nil)
		builder.On("CompileKey", jobSpecs[0].Name).Return(protoKey, nil)
		builder.On("CompileMessage", resource).Return(protoMsg, nil)
		defer builder.AssertExpectations(t)

		writerErr := errors.New("kafka is down")
		writer := new(mock.MetaWriter)
		writer.On("Write", protoKey, protoMsg).Return(writerErr)
		defer writer.AssertExpectations(t)

		service := meta.NewService(writer, builder)
		err := service.Publish(namespaceSpec, jobSpecs, nil)

		assert.NotNil(t, err)
		assert.Equal(t, "failed to write metadata message: job-1: kafka is down", err.Error())
	})
}
