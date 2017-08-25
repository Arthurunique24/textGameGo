package server

import (
	"../handlers"
	"net/http"
)


func RunHTTPServer(addr string) error {
	http.HandleFunc("/updateUData",handlers.UpdateUsersData)
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/signin", handlers.SignIn)
	http.HandleFunc("/start", handlers.Start)
	return http.ListenAndServe(addr, nil)
}
