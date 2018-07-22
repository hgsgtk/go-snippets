package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB init database connection
func InitDB(driver string, dsn string) error {
	var err error
	db, err = sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	if db.Ping() != nil {
		return err
	}
	return nil
}
