package libs

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/backent/go-article/helpers"
	_ "github.com/go-sql-driver/mysql"
)

func Initiate() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	DB_CONNECTION_MAX_LIFETIME_IN_SEC, err := strconv.Atoi(os.Getenv("DB_CONNECTION_MAX_LIFETIME_IN_SEC"))
	helpers.PanicIfError(err)
	DB_MAX_OPEN_CONNECTION, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTION"))
	helpers.PanicIfError(err)
	DB_MAX_IDLE_CONNECTION, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION"))
	helpers.PanicIfError(err)

	db.SetConnMaxLifetime(time.Second * time.Duration(DB_CONNECTION_MAX_LIFETIME_IN_SEC))
	db.SetMaxOpenConns(DB_MAX_OPEN_CONNECTION)
	db.SetMaxIdleConns(DB_MAX_IDLE_CONNECTION)

	return db
}
