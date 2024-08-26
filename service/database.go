package service

import (
	"database/sql"
	_ "github.com/lib/pq" 
	"97HW/config"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		panic(err)
	}
}

func getDB() *sql.DB {
	return db
}
