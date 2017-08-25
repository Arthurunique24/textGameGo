package handlers

import (
	"net/http"
	"../models"
)

func CheckSessionId(rw http.ResponseWriter, req *http.Request) {
	if models.CheckSessionId(req.Body) == true {
		rw.WriteHeader(http.StatusOK)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}
}
