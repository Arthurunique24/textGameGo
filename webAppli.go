package main

import (
    "./DAO"
    "./server"
    _ "github.com/lib/pq"
    "log"
)

func main() {
    DAO.Connect()
    log.Fatal(server.RunHTTPServer(":8080"))
    //http.HandleFunc("/signup", handlers.SignUp)
    //http.HandleFunc("/signin", handlers.SignIn)
    //log.Fatal(http.ListenAndServe(":8082", nil))
    defer DAO.Disconnect()
}
