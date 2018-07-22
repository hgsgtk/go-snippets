package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// NewDB create database global connection.
func NewDB(driver string, dsn string) (*sql.DB, error) {
	var err error
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if db.Ping() != nil {
		return nil, err
	}
	return db, nil
}
