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
	Db, err = sql.Open("mysql", "root:mysqlxjl12322@163.com@tcp(152.136.43.225:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
