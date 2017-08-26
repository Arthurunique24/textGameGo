package server

import (
	"net/http"

	"github.com/ChernovAndrey/textGameGo/handlers"
)

func RunHTTPServer(addr string) error {
	http.HandleFunc("/checkId", handlers.CheckSessionId)
	http.HandleFunc("/updateUData", handlers.UpdateUsersData)
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/signin", handlers.SignIn)
	http.HandleFunc("/start", handlers.Start)
	http.HandleFunc("/step", handlers.Step)
	return http.ListenAndServe(addr, nil)
}
