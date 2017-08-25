package models

import (
	"encoding/json"
	"io"
)


func Parse(p interface{}, body io.ReadCloser) error{
	decoder :=json.NewDecoder(body)
	defer body.Close()
	err:= decoder.Decode(&p)
	return err
}
