package stdlib

import (
	"database/sql"
	"database/sql/driver"
	"io"
)

func init() { sql.Register("pgx", drv{}) }

type drv struct{}

func (drv) Open(name string) (driver.Conn, error) { return conn{}, nil }

type conn struct{}

func (conn) Prepare(string) (driver.Stmt, error) { return stmt{}, nil }
func (conn) Close() error                        { return nil }
func (conn) Begin() (driver.Tx, error)           { return tx{}, nil }

type stmt struct{}

func (stmt) Close() error                               { return nil }
func (stmt) NumInput() int                              { return -1 }
func (stmt) Exec([]driver.Value) (driver.Result, error) { return result(1), nil }
func (stmt) Query([]driver.Value) (driver.Rows, error)  { return rows{}, nil }

type tx struct{}

func (tx) Commit() error   { return nil }
func (tx) Rollback() error { return nil }

type result int64

func (r result) LastInsertId() (int64, error) { return int64(r), nil }
func (r result) RowsAffected() (int64, error) { return int64(r), nil }

type rows struct{}

func (rows) Columns() []string { return []string{"id", "name", "email"} }
func (rows) Close() error      { return nil }
func (rows) Next(dest []driver.Value) error {
	if len(dest) >= 3 {
		dest[0] = "id"
		dest[1] = "name"
		dest[2] = "email"
	}
	return io.EOF
}
