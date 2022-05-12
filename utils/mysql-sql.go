package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
