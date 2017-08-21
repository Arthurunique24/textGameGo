package DAO

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "7680928160"
	DB_NAME     = "textgame"
)


var db *sql.DB

func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if(err!=nil){
		fmt.Println("fail")
		panic(err)
	}
}

func Disconnect() {
	 db.Close()
}