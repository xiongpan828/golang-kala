package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)
/*
初始化MySQL
*/
func InitMySQL() (err error) {
	dsn := "root:eaypassword@tcp(localhost:3306)/kala"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(20)
	return DB.Ping()
}
func Close()  {
	DB.Close()
}
