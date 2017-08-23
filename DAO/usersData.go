package DAO

import (
	"fmt"
	"errors"
)

func UpdateUsersData(userId int, rating int, countGames int) error{
	rows, err := db.Exec("Update  usersdata set (rating,countGames) = ($1, $2) where userId = $3",
		rating,countGames,userId)

	countUp,_:=rows.RowsAffected()
	if countUp == 0{
		err= errors.New("not found in usersData")
	}
	if err != nil {
		fmt.Println("UpdateUSersData",err)
	}
	return err
}
