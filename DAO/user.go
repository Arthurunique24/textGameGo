package DAO

import (
	"errors"
	"fmt"
)

func InsertUser(login, password string) error{
	_, err := db.Exec("INSERT INTO users (login,password) VALUES ($1, $2)",
		login,password)
	if err != nil {
		fmt.Println("insert User",err)
	}
	return err
}

func CheckUser(login,password string) (int,error){

	rows, err := db.Query("SELECT id FROM users where (login,password) = ($1,$2)",login,password) //id
	defer  rows.Close()
	if rows.Next()==false {
		err= errors.New("not found")
	}
	var id int
	rows.Scan(&id)
	return id,err
}

