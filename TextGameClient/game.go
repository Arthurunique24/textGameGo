package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"./SupportFunctions"
)

func signUp(url string, jsonStr []byte) string {
	//fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	//fmt.Println("Response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("Response Body:", string(body))
	return resp.Status
}

func signIn(url string, jsonStr []byte) ([]byte, string) {
	//fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	//fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("Response Body:", string(body))
	return body, resp.Status
}

func authorizeUser(url string) []byte {
	var login string
	var password string
	var sessionId []byte
	var signInStatus string

	fmt.Println("Hello, do you want to SignIn(1) or SignUp(2)?")
	var input string
	fmt.Scanf("%v\n", &input)

	if input == "SignIn" || input == "signIn" || input == "signin" || input == "1"{
		fmt.Print("Enter login: "); fmt.Scanf("%v", &login)
		fmt.Print("Enter password: "); fmt.Scanf("%v", &password)
		sessionId, signInStatus = signIn(url+"signin", []byte(SupportFunctions.TakeLoginAndPassword(login, password)))

		for signInStatus == "404 Not Found" {
			fmt.Println("Incorrect login or password, try again")
			fmt.Print("Enter login: "); fmt.Scanf("%v", &login)
			fmt.Print("Enter password: "); fmt.Scanf("%v", &password)
			sessionId, signInStatus = signIn(url+"signin", []byte(SupportFunctions.TakeLoginAndPassword(login, password)))
		}

	} else if input == "SignUp" || input == "signUp" || input == "signup" || input == "2"{
		fmt.Print("Enter login: "); fmt.Scanf("%v", &login)
		fmt.Print("Enter password: "); fmt.Scanf("%v", &password)
		signUpStatus := signUp(url+"signup", []byte(SupportFunctions.TakeLoginAndPassword(login, password)))

		for signUpStatus == "409 Conflict" {
			fmt.Println("Try another login or password")
			fmt.Print("Enter login: "); fmt.Scanf("%v", &login)
			fmt.Print("Enter password: "); fmt.Scanf("%v", &password)
			signUpStatus = signUp(url+"signup", []byte(SupportFunctions.TakeLoginAndPassword(login, password)))
		}
		sessionId, signInStatus = signIn(url+"signin", []byte(SupportFunctions.TakeLoginAndPassword(login, password)))

	} else {fmt.Println("I don't understand")}

	return sessionId
}

func Start(url string, idSession []byte) []byte {
	//fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(idSession))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("Start")
	//fmt.Println("Response Status:", resp.Status)
	//fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("Response Body:", string(body))

	return body
}

func Step(url string, stepJson []byte) []byte{
	//fmt.Println("URL:>", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(stepJson))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("Step")
	//fmt.Println("Response Status:", resp.Status)
	//fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("Response Body:", string(body))

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
	//fmt.Println("IdSession in main:", string(idJson))

	idSession, possibleSteps, finished, message =
		SupportFunctions.UnmarshalGameSession(Start(url+"start", idJson))
	fmt.Println("Possible steps:", possibleSteps)
	fmt.Println("Tip:", message)

	for finished != true {
		fmt.Println("Where do we go?")
		fmt.Scanf("%v", &userStep)
		idSession, possibleSteps, finished, message =
			SupportFunctions.UnmarshalGameSession(Step(url+"step", SupportFunctions.TakeGameStepJson(idSession, userStep)))
		fmt.Println("Possible steps:", possibleSteps)
		fmt.Println("Tip:", message)
	}
}