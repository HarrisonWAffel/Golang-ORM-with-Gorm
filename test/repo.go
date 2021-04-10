package test

import (
	"database/sql"
	"errors"
	"fmt"
	migrate "github.com/HarrisonWAffel/dbTrain/db"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"testing"
)

type BaseRepositoryTest interface {
	GetByIdTest(t *testing.T)
	CreateTest(t *testing.T)
	UpdateTest(t *testing.T)
	DeleteTest(t *testing.T)
}

//BaseRepoTester provides docker test suites
type BaseRepoTester struct {
	resource        *dockertest.Resource
	DSN             string
	NetworkSettings *docker.NetworkSettings
}

func (rt *BaseRepoTester) StartMockDatabase() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "10.5", []string{"POSTGRES_USER=postgres",
		"POSTGRES_DB=learn",
		"POSTGRES_PASSWORD=password123"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	var db *sql.DB
	rt.resource = resource
	rt.NetworkSettings = resource.Container.NetworkSettings
	resource.Expire(90)
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		resourceNetworkProperties := resource.Container.NetworkSettings.Ports["5432/tcp"][0]
		if resourceNetworkProperties.HostIP == "0.0.0.0" {
			resourceNetworkProperties.HostIP = "127.0.0.1"
		}

		rt.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			resourceNetworkProperties.HostIP, "postgres", "password123",
			"learn", resourceNetworkProperties.HostPort)

		var err error

		db, err = sql.Open("postgres", rt.DSN)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if db.Ping() == nil {
			migrateErr := migrate.AutoMigrate(rt.DSN)
			if migrateErr != nil {
				fmt.Println(migrateErr)
			} else {
				fmt.Println("Mock database migration success!")
			}
			return migrateErr
		}
		return errors.New("could not ping")
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
}

func (rt *BaseRepoTester) StopMockDatabase() {
	rt.resource.Close()
}

type BaseHandlerTest interface {
	GETTest(t *testing.T)
	POSTTest(t *testing.T)
	PUTTest(t *testing.T)
	DELETETest(t *testing.T)
}
