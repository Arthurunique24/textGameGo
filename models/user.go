package models

import (
	"io"
	"fmt"
	"net/http"
	"../DAO"
)
type User struct {
	Login string
	Password string
}


func (u *User) Insert(body io.ReadCloser) (int){
	//err:=u.Parse(body)
	 err :=Parse(u,body)
	if err !=nil {
		fmt.Println("parse json fail" , err) //err
		return http.StatusBadRequest
	}
	if DAO.InsertUser(u.Login, u.Password)!=nil {
		return http.StatusConflict
	}
	return http.StatusOK
}
func(u *User) Check(body io.ReadCloser) (int, interface{}){
	err := Parse(u,body)
	if err !=nil {
		fmt.Println("parse json fail", err)
		return http.StatusBadRequest,nil
	}
	var sessionId string
	sessionId,err=DAO.CheckUser(u.Login,u.Password)
	if  err != nil{
		return http.StatusNotFound,nil
	}
	return http.StatusOK,SessionId{Id:sessionId}
}
