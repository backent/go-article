package libs

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Initiate() *sql.DB {
	db, err := sql.Open("mysql", "root:adminlocal@tcp(localhost:3306)/articles")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
