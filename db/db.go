package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPGStorage(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db, nil
}
