package DAO

import (
	"fmt"
	"time"
)

//var DurationSession  = time.
func CheckId(id string) (bool){
	rows, err := db.Query("SELECT createDate FROM sessions where sessionid = $1",id)
	defer  rows.Close()
	if err !=nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() == false {
		return false
	}
	t := time.Now()
	rows.Scan(&t)
	fmt.Println(t)
	return true
}