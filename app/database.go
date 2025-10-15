package app

import (
	"belajar-rest-api-golang/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:alsya12345@tcp(localhost:3306)/golang_restful_api")
	helper.PanicIfError(err)

	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(time.Minute * 10)

	return db
}
