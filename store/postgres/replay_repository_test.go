// +build !unit_test

package postgres

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/odpf/optimus/job"
	"github.com/odpf/optimus/mock"
	"github.com/odpf/optimus/models"
	"github.com/stretchr/testify/assert"
)

func TestReplayRepository(t *testing.T) {
	projectSpec := models.ProjectSpec{
		ID:   uuid.Must(uuid.NewRandom()),
		Name: "t-optimus-id",
		Config: map[string]string{
			"bucket": "gs://some_folder",
		},
	}
	namespaceSpec := models.NamespaceSpec{
		ID:          uuid.Must(uuid.NewRandom()),
		Name:        "dev-team-1",
		ProjectSpec: projectSpec,
	}

	gTask := "g-task"
	jobConfigs := []models.JobSpec{
		{
			ID:   uuid.Must(uuid.NewRandom()),
			Name: "job-1",
		},
		{
			ID:   uuid.Must(uuid.NewRandom()),
			Name: "job-2",
		},
		{
			ID:   uuid.Must(uuid.NewRandom()),
			Name: "job-3",
		},
	}
	startTime, _ := time.Parse(job.ReplayDateFormat, "2020-01-15")
	endTime, _ := time.Parse(job.ReplayDateFormat, "2020-01-20")
	testConfigs := []*models.ReplaySpec{
		{
			ID:        uuid.Must(uuid.NewRandom()),
			StartDate: startTime,
			EndDate:   endTime,
			Status:    models.ReplayStatusAccepted,
		},
		{
			ID:        uuid.Must(uuid.NewRandom()),
			StartDate: startTime,
			EndDate:   endTime,
			Status:    models.ReplayStatusFailed,
		},
		{
			ID:        uuid.Must(uuid.NewRandom()),
			StartDate: startTime,
			EndDate:   endTime,
			Status:    models.ReplayStatusInProgress,
		},
	}

	DBSetup := func() *gorm.DB {
		dbURL, ok := os.LookupEnv("TEST_OPTIMUS_DB_URL")
		if !ok {
			panic("unable to find TEST_OPTIMUS_DB_URL env var")
		}
		dbConn, err := Connect(dbURL, 1, 1)
		if err != nil {
			panic(err)
		}
		m, err := NewHTTPFSMigrator(dbURL)
		if err != nil {
			panic(err)
		}
		if err := m.Drop(); err != nil {
			panic(err)
		}
		if err := Migrate(dbURL); err != nil {
			panic(err)
		}

		return dbConn
	}

	t.Run("Insert and GetByID", func(t *testing.T) {
		db := DBSetup()
		defer db.Close()

		execUnit1 := new(mock.TaskPlugin)
		defer execUnit1.AssertExpectations(t)

		for idx, jobConfig := range jobConfigs {
			jobConfig.Task = models.JobSpecTask{Unit: execUnit1}
			testConfigs[idx].Job = jobConfig
		}

		allTasksRepo := new(mock.SupportedTaskRepo)
		defer allTasksRepo.AssertExpectations(t)
		adapter := NewAdapter(allTasksRepo, nil)

		var testModels []*models.ReplaySpec
		testModels = append(testModels, testConfigs...)

		repo := NewReplayRepository(db, jobConfigs[0], adapter)
		err := repo.Insert(testModels[0])
		assert.Nil(t, err)

		checkModel, err := repo.GetByID(testModels[0].ID)
		assert.Nil(t, err)
		assert.Equal(t, testModels[0].ID, checkModel.ID)
	})

	t.Run("UpdateStatus", func(t *testing.T) {
		db := DBSetup()
		defer db.Close()
		var testModels []*models.ReplaySpec
		testModels = append(testModels, testConfigs...)

		execUnit1 := new(mock.TaskPlugin)
		defer execUnit1.AssertExpectations(t)

		for idx, jobConfig := range jobConfigs {
			jobConfig.Task = models.JobSpecTask{Unit: execUnit1}
			testConfigs[idx].Job = jobConfig
		}

		allTasksRepo := new(mock.SupportedTaskRepo)
		defer allTasksRepo.AssertExpectations(t)

		adapter := NewAdapter(allTasksRepo, nil)
		repo := NewReplayRepository(db, jobConfigs[0], adapter)
		err := repo.Insert(testModels[0])
		assert.Nil(t, err)

		errMessage := "failed to execute"
		replayMessage := models.ReplayMessage{
			Type:    "test failure",
			Message: errMessage,
		}
		err = repo.UpdateStatus(testModels[0].ID, models.ReplayStatusFailed, replayMessage)
		assert.Nil(t, err)

		checkModel, err := repo.GetByID(testModels[0].ID)
		assert.Nil(t, err)
		assert.Equal(t, models.ReplayStatusFailed, checkModel.Status)
		assert.Equal(t, errMessage, checkModel.Message.Message)
	})

	t.Run("GetJobByStatus", func(t *testing.T) {
		t.Run("should return list of job specs given list of status", func(t *testing.T) {
			db := DBSetup()
			defer db.Close()
			testModels := []*models.ReplaySpec{}
			testModels = append(testModels, testConfigs...)

			execUnit1 := new(mock.TaskPlugin)
			defer execUnit1.AssertExpectations(t)
			execUnit1.On("GetTaskSchema", context.TODO(), models.GetTaskSchemaRequest{}).Return(models.GetTaskSchemaResponse{
				Name: gTask,
			}, nil)
			for idx, jobConfig := range jobConfigs {
				jobConfig.Task = models.JobSpecTask{Unit: execUnit1}
				testConfigs[idx].Job = jobConfig
			}

			allTasksRepo := new(mock.SupportedTaskRepo)
			defer allTasksRepo.AssertExpectations(t)
			allTasksRepo.On("GetByName", gTask).Return(execUnit1, nil)
			adapter := NewAdapter(allTasksRepo, nil)

			unitData := models.GenerateTaskDestinationRequest{
				Config: models.TaskPluginConfigs{}.FromJobSpec(jobConfigs[0].Task.Config),
				Assets: models.TaskPluginAssets{}.FromJobSpec(jobConfigs[0].Assets),
			}
			execUnit1.On("GenerateTaskDestination", context.TODO(), unitData).Return(models.GenerateTaskDestinationResponse{Destination: "p.d.t"}, nil)

			projectJobSpecRepo := NewProjectJobSpecRepository(db, projectSpec, adapter)
			jobRepo := NewJobSpecRepository(db, namespaceSpec, projectJobSpecRepo, adapter)
			err := jobRepo.Insert(testModels[0].Job)
			assert.Nil(t, err)
			err = jobRepo.Insert(testModels[1].Job)
			assert.Nil(t, err)
			err = jobRepo.Insert(testModels[2].Job)
			assert.Nil(t, err)

			repo := NewReplayRepository(db, jobConfigs[0], adapter)
			err = repo.Insert(testModels[0])
			assert.Nil(t, err)
			err = repo.Insert(testModels[1])
			assert.Nil(t, err)
			err = repo.Insert(testModels[2])
			assert.Nil(t, err)

			statusList := []string{models.ReplayStatusAccepted, models.ReplayStatusInProgress}
			replays, err := repo.GetByStatus(statusList)
			assert.Nil(t, err)
			assert.Equal(t, jobConfigs[0].ID, replays[0].Job.ID)
			assert.Equal(t, jobConfigs[2].ID, replays[1].Job.ID)
		})
	})

	t.Run("GetJobByIDAndStatus", func(t *testing.T) {
		t.Run("should return list of job specs given job_id and list of status", func(t *testing.T) {
			db := DBSetup()
			defer db.Close()
			var testModels []*models.ReplaySpec
			testModels = append(testModels, testConfigs...)

			execUnit1 := new(mock.TaskPlugin)
			defer execUnit1.AssertExpectations(t)
			execUnit1.On("GetTaskSchema", context.TODO(), models.GetTaskSchemaRequest{}).Return(models.GetTaskSchemaResponse{
				Name: gTask,
			}, nil)
			for idx, jobConfig := range jobConfigs {
				jobConfig.Task = models.JobSpecTask{Unit: execUnit1}
				testConfigs[idx].Job = jobConfig
			}

			allTasksRepo := new(mock.SupportedTaskRepo)
			defer allTasksRepo.AssertExpectations(t)
			allTasksRepo.On("GetByName", gTask).Return(execUnit1, nil)
			adapter := NewAdapter(allTasksRepo, nil)

			unitData := models.GenerateTaskDestinationRequest{
				Config: models.TaskPluginConfigs{}.FromJobSpec(jobConfigs[0].Task.Config),
				Assets: models.TaskPluginAssets{}.FromJobSpec(jobConfigs[0].Assets),
			}
			execUnit1.On("GenerateTaskDestination", context.TODO(), unitData).Return(models.GenerateTaskDestinationResponse{Destination: "p.d.t"}, nil)

			projectJobSpecRepo := NewProjectJobSpecRepository(db, projectSpec, adapter)
			jobRepo := NewJobSpecRepository(db, namespaceSpec, projectJobSpecRepo, adapter)
			err := jobRepo.Insert(testModels[0].Job)
			assert.Nil(t, err)
			err = jobRepo.Insert(testModels[1].Job)
			assert.Nil(t, err)
			err = jobRepo.Insert(testModels[2].Job)
			assert.Nil(t, err)

			repo := NewReplayRepository(db, jobConfigs[0], adapter)
			err = repo.Insert(testModels[0])
			assert.Nil(t, err)
			err = repo.Insert(testModels[1])
			assert.Nil(t, err)
			err = repo.Insert(testModels[2])
			assert.Nil(t, err)

			statusList := []string{models.ReplayStatusAccepted, models.ReplayStatusInProgress}
			replays, err := repo.GetByJobIDAndStatus(testModels[2].Job.ID, statusList)
			assert.Nil(t, err)
			assert.Equal(t, jobConfigs[2].ID, replays[0].Job.ID)
		})
	})
}
