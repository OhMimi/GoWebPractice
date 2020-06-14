package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:X1nt0neo924###@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
}
