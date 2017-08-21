package models

import (
	"io"
	"encoding/json"
	"fmt"
	"net/http"
	"../DAO"
)
type User struct {
	Login string
	Password string
}

func (u *User) Parse(body io.ReadCloser) error{
	decoder :=json.NewDecoder(body)
	defer body.Close()
	err:= decoder.Decode(&u)
	return err
}

func (u *User) Insert(body io.ReadCloser) int{
	if u.Parse(body) !=nil {
		fmt.Println("parse json fail")
		return http.StatusBadRequest
	}
	fmt.Println("user insert")
	if DAO.InsertUser(u.Login, u.Password)!=nil {
		return http.StatusConflict
	}
	return http.StatusOK
}
func(u *User) Check(body io.ReadCloser) int{
	if u.Parse(body) !=nil {
		fmt.Println("parse json fail")
		return http.StatusBadRequest
	}

	if  DAO.CheckUser(u.Login,u.Password)!=nil{
		return http.StatusNotFound
	}
	return http.StatusOK
}
