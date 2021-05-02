package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"adi123",
		"127.0.0.1:3306",
		"users_db",
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
