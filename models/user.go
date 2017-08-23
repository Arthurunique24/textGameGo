package models

import (
	"io"
	"fmt"
	"net/http"
	"../DAO"
	"../session"
)
type User struct {
	Login string
	Password string
}

/*func (u *User) Parse(body io.ReadCloser) error{
	decoder :=json.NewDecoder(body)
	defer body.Close()
	err:= decoder.Decode(&u)
	return err

}*/

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
	var id int
	id,err=DAO.CheckUser(u.Login,u.Password)
	if  err != nil{
		return http.StatusNotFound,nil
	}
	if session.CheckId(id)==false {
		session.Create(id)
	}
	return http.StatusOK,Id{Id:id}
}
