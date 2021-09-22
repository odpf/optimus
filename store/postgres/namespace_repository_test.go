// +build !unit_test

package postgres

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/odpf/optimus/models"
	"github.com/stretchr/testify/assert"
)

func TestNamespaceRepository(t *testing.T) {
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

	transporterKafkaBrokerKey := "KAFKA_BROKERS"
	hash, _ := models.NewApplicationSecret("32charshtesthashtesthashtesthash")

	secrets := []models.ProjectSecretItem{
		{
			ID:    uuid.Must(uuid.NewRandom()),
			Name:  "g-optimus",
			Value: "secret",
		},
		{
			ID:    uuid.Must(uuid.NewRandom()),
			Name:  "t-optimus",
			Value: "super-secret",
		},
	}
	projectSpec := models.ProjectSpec{
		ID:   uuid.Must(uuid.NewRandom()),
		Name: "t-optimus",
		Config: map[string]string{
			"bucket":                  "gs://some_folder",
			transporterKafkaBrokerKey: "10.12.12.12:6668,10.12.12.13:6668",
		},
	}
	namespaceSpecs := []models.NamespaceSpec{
		{
			ID:   uuid.Must(uuid.NewRandom()),
			Name: "g-optimus",
		},
		{
			Name: "",
		},
		{
			ID:   uuid.Must(uuid.NewRandom()),
			Name: "t-optimus",
			Config: map[string]string{
				"bucket":                  "gs://some_folder",
				transporterKafkaBrokerKey: "10.12.12.12:6668,10.12.12.13:6668",
			},
		},
		{
			ID:   uuid.Must(uuid.NewRandom()),
			Name: "t-optimus-2",
			Config: map[string]string{
				"bucket":                  "gs://some_folder-2",
				transporterKafkaBrokerKey: "10.12.12.12:6668,10.12.12.13:6668",
			},
		},
	}

	t.Run("Insert", func(t *testing.T) {
		db := DBSetup()
		defer db.Close()
		testModels := []models.NamespaceSpec{}
		testModels = append(testModels, namespaceSpecs...)

		// save project
		projRepo := NewProjectRepository(db, hash)
		err := projRepo.Save(projectSpec)
		assert.Nil(t, err)

		secretRepo := NewSecretRepository(db, projectSpec, hash)
		err = secretRepo.Insert(secrets[0])
		assert.Nil(t, err)
		err = secretRepo.Insert(secrets[1])
		assert.Nil(t, err)

		repo := NewNamespaceRepository(db, projectSpec, hash)

		err = repo.Insert(testModels[0])
		assert.Nil(t, err)

		err = repo.Insert(testModels[1])
		assert.NotNil(t, err)

		checkModel, err := repo.GetByName(testModels[0].Name)
		assert.Nil(t, err)
		assert.Equal(t, "g-optimus", checkModel.Name)
		assert.Equal(t, projectSpec.Name, checkModel.ProjectSpec.Name)
		assert.Equal(t, 2, len(checkModel.ProjectSpec.Secret))
	})

	t.Run("Upsert", func(t *testing.T) {
		t.Run("insert different resource should insert two", func(t *testing.T) {
			db := DBSetup()
			defer db.Close()
			testModelA := namespaceSpecs[0]
			testModelB := namespaceSpecs[2]

			repo := NewNamespaceRepository(db, projectSpec, hash)

			//try for create
			err := repo.Save(testModelA)
			assert.Nil(t, err)

			checkModel, err := repo.GetByName(testModelA.Name)
			assert.Nil(t, err)
			assert.Equal(t, "g-optimus", checkModel.Name)

			//try for update
			err = repo.Save(testModelB)
			assert.Nil(t, err)

			checkModel, err = repo.GetByName(testModelB.Name)
			assert.Nil(t, err)
			assert.Equal(t, "t-optimus", checkModel.Name)
			assert.Equal(t, "10.12.12.12:6668,10.12.12.13:6668", checkModel.Config[transporterKafkaBrokerKey])
		})
		t.Run("insert same resource twice should overwrite existing", func(t *testing.T) {
			db := DBSetup()
			defer db.Close()
			testModelA := namespaceSpecs[2]

			repo := NewNamespaceRepository(db, projectSpec, hash)

			//try for create
			testModelA.Config["bucket"] = "gs://some_folder"
			err := repo.Save(testModelA)
			assert.Nil(t, err)

			checkModel, err := repo.GetByName(testModelA.Name)
			assert.Nil(t, err)
			assert.Equal(t, "t-optimus", checkModel.Name)

			//try for update
			testModelA.Config["bucket"] = "gs://another_folder"
			err = repo.Save(testModelA)
			assert.Nil(t, err)

			checkModel, err = repo.GetByName(testModelA.Name)
			assert.Nil(t, err)
			assert.Equal(t, "gs://another_folder", checkModel.Config["bucket"])
		})
		t.Run("upsert without ID should auto generate it", func(t *testing.T) {
			db := DBSetup()
			defer db.Close()
			testModelA := namespaceSpecs[0]
			testModelA.ID = uuid.Nil

			repo := NewNamespaceRepository(db, projectSpec, hash)

			//try for create
			err := repo.Save(testModelA)
			assert.Nil(t, err)

			checkModel, err := repo.GetByName(testModelA.Name)
			assert.Nil(t, err)
			assert.Equal(t, "g-optimus", checkModel.Name)
			assert.Equal(t, 36, len(checkModel.ID.String()))
		})
	})

	t.Run("GetByName", func(t *testing.T) {
		db := DBSetup()
		defer db.Close()
		testModels := []models.NamespaceSpec{}
		testModels = append(testModels, namespaceSpecs...)

		repo := NewNamespaceRepository(db, projectSpec, hash)

		err := repo.Insert(testModels[0])
		assert.Nil(t, err)

		checkModel, err := repo.GetByName(testModels[0].Name)
		assert.Nil(t, err)
		assert.Equal(t, "g-optimus", checkModel.Name)
	})

	t.Run("GetAll", func(t *testing.T) {
		db := DBSetup()
		defer db.Close()
		testModels := []models.NamespaceSpec{}
		testModels = append(testModels, namespaceSpecs...)

		repo := NewNamespaceRepository(db, projectSpec, hash)

		err := repo.Insert(testModels[0])
		assert.Nil(t, err)
		err = repo.Insert(testModels[2])
		assert.Nil(t, err)

		checkModel, err := repo.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(checkModel))
	})
}
