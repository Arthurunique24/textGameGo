package DAO

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "textgame"
)

var db *sql.DB

func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("fail")
		panic(err)
	}
}

func Disconnect() {
	db.Close()
}
