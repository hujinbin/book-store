package utils

import(
	"database/sql"
)

var (
	Db * sql.DB
	err error
)

func init(){
	Db, err = sql.Open("mysql", "bookstore:Bookstore.com@_00..@tcp(62.234.11.1:3306)/bookstore")
	if err != nil{
      panic(err.Error())
	}
}