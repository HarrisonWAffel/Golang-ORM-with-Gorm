package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
	"path/filepath"
	"runtime"
)

func AutoMigrate(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	_, b, _, _ := runtime.Caller(0)
	//get path to this specific go file.
	//Allows others to call this function
	//without worrying about caller file placement
	basepath := filepath.Dir(b)

	//do a quick check on the basepath to see if its a valid directory. Base path should
	// be valid when running the service locally, and invalid when deployed in a container.
	// This is due to the lack of a GOPATH within the final stage of container.
	_, err = os.Stat(basepath + "/setup.go")
	if err != nil {
		basepath = "/db"
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/migrations", basepath),
		"postgres", driver)
	if err != nil {
		return err
	}

	return m.Up()
}
