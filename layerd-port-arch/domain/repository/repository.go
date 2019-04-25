package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	Exec interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
	}
	Query interface {
		Query(query string, args ...interface{}) (*sql.Rows, error)
		Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
		QueryRowx(query string, args ...interface{}) *sqlx.Row
	}
	Prepare interface {
		Preparex(query string) (*sqlx.Stmt, error)
	}
	Begin interface {
		Beginx() (*sqlx.Tx, error)
	}
)

type DBHandler interface {
	Exec
	Query
	Prepare
}

type DBConnector interface {
	DBHandler
	Begin
}
