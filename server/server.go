package server

import (
	"time"
	"net/http"
	"io"
	"fmt"
	"encoding/json"
	"../server/workers"
	"../game"
	"../handlers"
)

type gameSession struct {
	Status string `json:"status,omitempty"`
	Body struct {
		Turn int `json:"turn,omitempty"`
		Answer []int `json:"answer,omitempty"`
	} `json:"body,omitempty"`
}

var wp *workers.Pool = workers.NewPool(5)

func init() {
	wp.Run()
}

const requestWaitInQueueTimeout = time.Millisecond * 100

// ab -c10 -n20 localhost:8000/hello

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := wp.AddTaskSyncTimed(func() interface{} {
		time.Sleep(time.Millisecond * 500)
		io.WriteString(w, "Hello, world!")
		return nil
	}, requestWaitInQueueTimeout)



	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}
}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	_, err := wp.AddTaskSyncTimed(func() interface{} {
		var gs gameSession
		gs.parse(req.Body)
		var gsAnswer gameSession
		gsAnswer.Status = "Ok"
		gsAnswer.Body.Turn = gs.Body.Turn
		_, gsAnswer.Body.Answer = game.Start()
		err := json.NewEncoder(rw).Encode(gsAnswer)
		return err
	}, requestWaitInQueueTimeout)



	if err != nil {
		//http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}
}
func (session *gameSession) parse(body io.ReadCloser) error {
	decoder :=json.NewDecoder(body)
	defer body.Close()
	err:= decoder.Decode(&session)
	fmt.Println(session.Status)
	return err
}

func RunHTTPServer(addr string) error {
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/signin", handlers.SignIn)
	http.HandleFunc("/hello", rootHandler)
	http.HandleFunc("/test", testHandler)
	return http.ListenAndServe(addr, nil)
}
