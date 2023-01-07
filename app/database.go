package app

import (
	"database/sql"
	"golang-restapi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_spicesignature?parseTime=true")
	helper.PanicIfError(err)

	//Database Pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
