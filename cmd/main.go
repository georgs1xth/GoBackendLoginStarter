package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/georgs1xth/APIBACKEND/cmd/api"
	"github.com/georgs1xth/APIBACKEND/config"
	"github.com/georgs1xth/APIBACKEND/db"
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

	initStorage(db)

	defer db.Close()
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: successfully connected")
}
