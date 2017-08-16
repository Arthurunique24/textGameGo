package User

import (
	"io"
	"fmt"
	"encoding/json"
)
type User struct {
	Login string
	Password string
}

func (u *User) Parse(body io.ReadCloser) error{
	fmt.Println("hello")
	decoder :=json.NewDecoder(body)
	defer body.Close()
	err:= decoder.Decode(&u)
	//fmt.Println(err.Error())
	return err
}