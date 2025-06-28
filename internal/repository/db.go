package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
    "log"
)

var db *sql.DB

func InitPostgres(dsn string) {
	var err error
	db, err = sql.Open("postgre", dsn)
	if err != nil {
		log.Fatal(err)
	}
}