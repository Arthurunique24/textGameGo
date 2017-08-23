package models

import (
	"fmt"
	"net/http"
	"io"
	"../DAO"
)

type UsersData struct{
	 UserId int
	 Rating int
	 CountGames int
}

func (ud *UsersData) Update(body io.ReadCloser) int{
	err :=Parse(ud,body)
	if err !=nil {
		fmt.Println("parse json fail" , err)
		return http.StatusBadRequest
	}
	if DAO.UpdateUsersData(ud.UserId, ud.Rating,ud.CountGames)!=nil {
		return http.StatusNotFound
	}
	return http.StatusOK
}