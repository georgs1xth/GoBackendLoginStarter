package main

import (
	"fmt"
	"log"
	"os"

	"github.com/georgs1xth/APIBACKEND/config"
	"github.com/georgs1xth/APIBACKEND/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf(`user=%s 
							password=%s 
							dbname=%s 
							sslmode=disable`,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
	)
	db, err := db.NewPGStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"BackendAPI",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
