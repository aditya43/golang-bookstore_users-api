package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aditya43/golang-bookstore_users-api/utils/env"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBClient *sql.DB

	db_username = env.Get("MYSQL_DB_USERNAME")
	db_password = env.Get("MYSQL_DB_PASSWORD")
	db_host     = env.Get("MYSQL_DB_HOST")
	db_schema   = env.Get("MYSQL_DB_SCHEMA")
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		db_username,
		db_password,
		db_host,
		db_schema,
	)

	var err error
	DBClient, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = DBClient.Ping(); err != nil {
		panic(err)
	}

	log.Println("Successfully connected to users_db database")
}
