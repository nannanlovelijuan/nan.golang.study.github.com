package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func init() {
	DB, _ = sql.Open("mysql", "root:123456@tcp(192.168.12.118:3306)/ezp-bigdata?charset=utf8mb4")
}

func Ping() error {
	return DB.Ping()
}

//func Query()  {
//	rows := DB.QueryRow("SELECT * FROM ezp_bd_easybi_db_source")
//}


