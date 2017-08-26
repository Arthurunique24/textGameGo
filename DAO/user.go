package DAO

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func InsertUser(login, password string) error {
	hash, err := hashPassword(password)
	if err != nil {
		fmt.Println("problem with hash", err)
	}
	_, err = db.Exec("INSERT INTO users (login,password) VALUES ($1, $2)",
		login, hash)
	if err != nil {
		fmt.Println("insert User", err)
	}
	return err
}

func CheckUser(login, password string) (string, error) {
	rows, err := db.Query("SELECT id,password FROM users where login = $1", login) //id
	defer rows.Close()
	if rows.Next() == false {
		err = errors.New("not found")
	}
	var id int
	var hash string
	rows.Scan(&id, &hash)
	if checkPasswordHash(password, hash) == false {
		err = errors.New("not found")
	}
	return createSession(id), err
}

func createSession(id int) string {
	sessionID, err := generateRandomString(32)
	if err != nil {
		fmt.Println("create sessions fail", err)
		return ""
	}
	_, err = db.Exec("INSERT INTO sessions (userId,sessionId) VALUES ($1, $2) ON CONFLICT (userId) DO UPDATE "+
		"SET sessionId = excluded.SessionId;",
		id, sessionID)
	if err != nil {
		fmt.Println("insert session", err)
	}
	return sessionID
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
