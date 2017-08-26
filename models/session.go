package models

import (
	"fmt"
	"io"

	"github.com/ChernovAndrey/textGameGo/DAO"
)

func CheckSessionId(body io.ReadCloser) bool {
	fmt.Println("alp")
	var id = new(SessionId)
	err := Parse(id, body)
	if err != nil {
		fmt.Println("parse json fail", err)
		return false
	}
	fmt.Println("alp")
	fmt.Println(id.Id)
	return DAO.CheckId(id.Id)
}
