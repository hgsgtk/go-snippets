package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Datastore is interface of database global connection.
type Datastore interface {
	GetAllBooks() ([]*Book, error)
}

// DB is struct of sql.DB.
type DB struct {
	*sql.DB
}

// NewDB create datastore interface.
func NewDB(driver string, dsn string) (*DB, error) {
	var err error
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if db.Ping() != nil {
		return nil, err
	}
	return &DB{db}, nil
}
