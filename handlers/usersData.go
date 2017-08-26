package handlers

import (
	"net/http"

	"github.com/ChernovAndrey/textGameGo/models"
)

func UpdateUsersData(rw http.ResponseWriter, req *http.Request) {
	var ud models.UsersData
	rw.WriteHeader(ud.Update(req.Body))
}
