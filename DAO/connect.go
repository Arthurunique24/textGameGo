package DAO

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

/*const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "textgame"
)*/
type ConfigParameter struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

}
var config ConfigParameter

func init() {
	filename, _ := filepath.Abs("./config/database.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("config DB",err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("config DB",err)
		}
}

var db *sql.DB

func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("fail")
		panic(err)
	}
}

func Disconnect() {
	db.Close()
}
