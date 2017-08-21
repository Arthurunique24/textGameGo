package DAO

import (
	"errors"
	"fmt"
)

func InsertUser(login, password string) error{

	_, err := db.Exec("INSERT INTO users (login,password) VALUES ($1, $2)",
		login,password)
	fmt.Println(err)
	return err
}

func CheckUser(login,password string) error{

	rows, err := db.Query("SELECT * FROM users where (login,password) = ($1,$2)",login,password)
	defer  rows.Close()
	if rows.Next()==false {
		err= errors.New("not found")
	}
	return err
}

