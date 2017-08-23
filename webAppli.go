package main

import (
	"log"
	"./DAO"
	"./server"
)


func main() {
	DAO.Connect()
	log.Fatal(server.RunHTTPServer(":8080"))
	//http.HandleFunc("/signup", handlers.SignUp)
	//http.HandleFunc("/signin", handlers.SignIn)
	//log.Fatal(http.ListenAndServe(":8082", nil))
	defer DAO.Disconnect()
}