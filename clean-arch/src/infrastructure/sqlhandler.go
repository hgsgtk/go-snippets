package infrastructure

import (
	"database/sql"

	"github.com/Khigashiguchi/go-snippets/clean-arch/src/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/simple")
	if err != nil {
		panic(err.Error())
	}
	sh := new(SqlHandler)
	sh.Conn = conn
	return sh
}

func (handler *SqlHandler) Execute(smt string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(smt, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(smt string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Conn.Query(smt, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return res, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
