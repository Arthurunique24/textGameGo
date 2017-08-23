package models

import (
	"encoding/json"
	"io"
)

type Parsing interface {}

func Parse(p Parsing, body io.ReadCloser) error{
	decoder :=json.NewDecoder(body)
	defer body.Close()
	err:= decoder.Decode(&p)
	return err
}
