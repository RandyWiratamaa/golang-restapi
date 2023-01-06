package main

import (
	"database/sql"
	"log"
	"time"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/rest-go?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	//Database Pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
