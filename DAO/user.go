package DAO

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(login, password string) error{
	hash,err:=hashPassword(password)
	if err !=nil {
		fmt.Println("problem with hash", err)
	}
	_, err = db.Exec("INSERT INTO users (login,password) VALUES ($1, $2)",
		login,hash)
	if err != nil {
		fmt.Println("insert User",err)
	}

	return err
}

func CheckUser(login,password string) (int,error){
	rows, err := db.Query("SELECT id,password FROM users where login = $1",login) //id
	defer  rows.Close()
	if rows.Next()==false {
		err= errors.New("not found")
	}
	var id int
	var hash string
	rows.Scan(&id,&hash)
	if checkPasswordHash(password,hash) == false{
		err= errors.New("not found")
	}
	return id,err
}


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

