package Old

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

type GameSession struct {
	Status string `json:"status,omitempty"`
	Body struct {
		Turn string
	} `json:"body,omitempty"`
}

type GameAnswer struct {
	Status string `json:"status,omitempty"`
	Body struct {
		Answer []int `json:"answer"`
	} `json:"body,omitempty"`
}

func takeJSON(position string, turn string) string {
	gameSession := &GameSession{Status: position, Body: struct{ Turn string }{Turn: turn}}
	marshaled, err := json.Marshal(gameSession)
	if err != nil {
		fmt.Println(err)
		return "Error"
	}
	//fmt.Println(string(marshaled))
	return string(marshaled)
}

func postRequest(url string, jsonStr []byte) {
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

	fmt.Println("From unmarshal:", unmarshal(body))
}

func unmarshal(jsonStr []byte) []int {
	//jsonSts := []byte(`{"status":"start", "body": {"answer":[4, 3, 2]}}`)
	var gm GameAnswer
	//err := json.Unmarshal(jsonSts, &gm)

	if err := json.Unmarshal(jsonStr, &gm); err != nil {
		panic(err)
	}
	ints := gm.Body.Answer
	fmt.Println("Status: ", gm.Status)
	//fmt.Println("From unmarshal: ", ints)
	return ints
}

func main()  {
	url := "http://localhost:8080/test"
	jsonStr := takeJSON("start", "0")
	postRequest(url, []byte(jsonStr))
}
