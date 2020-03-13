package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

// 初始化
func Init(dns string) (err error) {

	DB,err = sqlx.Open("mysql",dns)
	if err != nil {
		return err
	}
	// 查看是否连成功
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
// Init 初始化
//func Init(mysqlDSN string) {
//	db = sqlx.MustConnect("mysql", mysqlDSN)
//	db.SetMaxIdleConns(1)
//	db.SetMaxOpenConns(3)
//}