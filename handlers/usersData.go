package handlers

import (
	"net/http"
	"../models"
)

func UpdateUsersData(rw http.ResponseWriter, req *http.Request){
	var ud models.UsersData
	rw.WriteHeader(ud.Update(req.Body))
}
