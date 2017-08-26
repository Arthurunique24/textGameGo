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
	Finished bool `json:"finished,omitempty"`
	Message string `json:"message,omitempty"`
}

type GameStep struct {
	Id string `json:"id,omitempty"`
	Step int `json:"step,omitempty"`
}

func unmarshalGameSession(jsonStr []byte) (string, []int, bool, string) {
	var gameSession GameSession

	if err := json.Unmarshal(jsonStr, &gameSession); err != nil {
		panic(err)
	}
	idSession := gameSession.Id
	possibleSteps := gameSession.PossibleSteps
	finished := gameSession.Finished
	message := gameSession.Message
	//fmt.Println("From unmarshal: ", idSession)
	return idSession, possibleSteps, finished, message
}

func takeLoginAndPassword(login string, password string) []byte {
	authSession := &AuthSession{Login: login, Password: password}
	marshaled, err := json.Marshal(authSession)
	if err != nil {
		fmt.Println(err)
	}
	return marshaled
}

func takeGameStepJson(id string, step int) []byte {
	gameStep := &GameStep{Id: id, Step: step}
	marshaled, err := json.Marshal(gameStep)
	if err != nil {
		fmt.Println(err)
	}
	return marshaled
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

func Start(url string, idSession []byte) []byte {
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(idSession))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Start")
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))

	return body
}

func Step(url string, stepJson []byte) []byte{
	fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(stepJson))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Start")
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))

	return body
}

func main() {
	url := "http://localhost:8080/"
	var idJson []byte
	var idSession string
	var possibleSteps []int
	var finished bool = false
	var message string
	var userStep int

	idJson = authorizeUser(url)
	fmt.Println("IdSession in main:", string(idJson))


	idSession, possibleSteps, finished, message = unmarshalGameSession(Start(url+"start", idJson))
	fmt.Println("Possible steps:", possibleSteps)
	fmt.Println("Tip:", message)

	for finished != true {
		fmt.Println("Where do we go?")
		fmt.Scanf("%v", &userStep)
		idSession, possibleSteps, finished, message = unmarshalGameSession(Step(url+"step", takeGameStepJson(idSession, userStep)))
		fmt.Println("Possible steps:", possibleSteps)
		fmt.Println("Tip:", message)
	}
	if finished == true {
		fmt.Println("Our congratulations! You won the game!")
	}

}