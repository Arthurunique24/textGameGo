package DAO

import "fmt"

func CheckId(id string) bool {
	rows, err := db.Query("SELECT userId FROM sessions where sessionid = $1", id)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	return (rows.Next()) && (err == nil)

}
