package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ChernovAndrey/textGameGo/DAO"
	"github.com/ChernovAndrey/textGameGo/server"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port string
}

var config Config

func init() {
	filename, _ := filepath.Abs("./config/server.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("config Port", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("config Port", err)
	}

}

func main() {
	DAO.Connect()
	log.Fatal(server.RunHTTPServer(config.Port))
	defer DAO.Disconnect()
}
