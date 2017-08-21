package handlers

import (
	"../models"
	"net/http"
	_ "github.com/lib/pq"
)

func SignUp(rw http.ResponseWriter, req *http.Request) {
	var u models.User
	rw.WriteHeader(u.Insert(req.Body))
}



func SignIn(rw http.ResponseWriter, req *http.Request) {
	var u models.User
	rw.WriteHeader(u.Check(req.Body))
}