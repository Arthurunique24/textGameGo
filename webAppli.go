package main

import (
	"./User"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func signUpHandler(rw http.ResponseWriter, req *http.Request) {
	var u User.User
	if u.Parse(req.Body) !=nil {
		fmt.Println("parse json fail")
		return
	}
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if(err!=nil){
		fmt.Println("fail")
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println(u.Login)
	_, err = db.Exec("INSERT INTO users (login,password) VALUES ($1, $2)",
		u.Login,u.Password )
	if (err!=nil){
		fmt.Println(err.Error())
		rw.WriteHeader(http.StatusConflict)
		return
	}
	rw.WriteHeader(http.StatusOK)
}



func signInHandler(rw http.ResponseWriter, req *http.Request) {
	var u User.User
	if u.Parse(req.Body) !=nil {
		fmt.Println("parse json fail")
		return
	}
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if(err!=nil){
		fmt.Println("fail")
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where (login,password) = ($1,$2)",u.Login,u.Password)
	defer  rows.Close()
	if (err!=nil)||(rows.Next()==false) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "7680928160"
	DB_NAME     = "textgame"
)


func main() {
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/signin", signInHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}