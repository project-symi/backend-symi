package infrastructure

import (
	"database/sql"
	"os"

	"project-symi-backend/app/interfaces/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
	godotenv.Load()
	conn, err := sql.Open("mysql", os.Getenv("DB_INFO"))
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

func (handler *SqlHandler) Begin() (database.Tx, error) {
	beginTx, err := handler.Conn.Begin()
	if err != nil {
		panic(err)
	}
	tx := new(SqlTx)
	tx.Tx = beginTx
	return tx, nil
}

type SqlTx struct {
	Tx *sql.Tx
}

func (tx *SqlTx) Exec(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := tx.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (tx *SqlTx) Rollback() (err error) {
	err = tx.Rollback()
	return
}

func (tx *SqlTx) Commit() (err error) {
	err = tx.Commit()
	return
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
