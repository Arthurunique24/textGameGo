package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

type AuthSession struct {
	Login string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

type GameSession struct {
	Id string `json:"id,omitempty"`
	PossibleSteps []int `json:"possibleSteps,omitempty"`
}

func unmarshal(jsonStr []byte) string {
	var gameSession GameSession

	if err := json.Unmarshal(jsonStr, &gameSession); err != nil {
		panic(err)
	}
	idSession := gameSession.Id
	//fmt.Println("From unmarshal: ", idSession)
	return idSession
}

func takeLoginAndPassword(login string, password string) string {
	authSession := &AuthSession{Login: login, Password: password}
	marshaled, err := json.Marshal(authSession)
	if err != nil {
		fmt.Println(err)
		return "Error in takeLoginAndPassword!"
	}
	//fmt.Println(string(marshaled))
	return string(marshaled)
}

func signIn(url string, jsonStr []byte) []byte{
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))

	return body
}

func signUp(url string, jsonStr []byte) {
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

func authorizeUser(url string) []byte {
	var login string
	var password string
	var sessionId []byte

	fmt.Println("Hello, do you want to SignIn(1) or SignUp(2)?")
	var input string
	fmt.Scanf("%v\n", &input)

	if input == "SignIn" || input == "signIn" || input == "signin" || input == "1"{
		fmt.Print("Enter login: "); fmt.Scanf("%v", &login)
		fmt.Print("Enter password: "); fmt.Scanf("%v", &password)
		sessionId = signIn(url+"signin", []byte(takeLoginAndPassword(login, password)))

	} else if input == "SignUp" || input == "signUp" || input == "signup" || input == "2"{
		fmt.Print("Enter login: "); fmt.Scanf("%v", &login)
		fmt.Print("Enter password: "); fmt.Scanf("%v", &password)
		signUp(url+"signup", []byte(takeLoginAndPassword(login, password)))
		sessionId = signIn(url+"signin", []byte(takeLoginAndPassword(login, password)))

	} else {fmt.Println("I don't understand")}

	return sessionId
}

func main()  {
	//var jsonStr = []byte(`{"login":"kek1","password":"1234"}`)
	url := "http://localhost:8080/"

	marhJson := authorizeUser(url)
	fmt.Println(string(marhJson))
}