package handlers

import (
	"../models"
	"net/http"
	"encoding/json"
	"fmt"
)

func SignUp(rw http.ResponseWriter, req *http.Request) {
	var u models.User
	rw.WriteHeader(u.Insert(req.Body))
}



func SignIn(rw http.ResponseWriter, req *http.Request) {
	var u models.User
	status,response := u.Check(req.Body)
	rw.WriteHeader(status)
	if response != nil {
		rw.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(rw).Encode(response)
		if err != nil {
			fmt.Println("incorrect format response", err)
		}
	}
}
