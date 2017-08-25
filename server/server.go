package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ChernovAndrey/textGameGo/game"
	"github.com/ChernovAndrey/textGameGo/handlers"
	"github.com/ChernovAndrey/textGameGo/server/workers"
)

type networkData struct {
	Status string `json:"status,omitempty"`
	Body   struct {
		Turn    int   `json:"turn,omitempty"`
		Answer  []int `json:"answer,omitempty"`
		Message string `json:"message,omitempty"`
    Finished bool `json:"finished,omitempty"`
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
		var inGS networkData
		inGS.parse(req.Body)
		var outGS networkData
		if inGS.Status == "Start" {
			outGS.Status = 200
			outGS.Body.Message = "Игра началась"
			outGS.Body.Answer = game.Start()
      outGS.Body.Finished = false
		} else if inGS.Status == "Step" {
			correct, finished, message, answer := game.Turn()
      outGS.Status = correct ? 200 : 500
      outGS.Body.Answer = answer
      outGS.Body.Message = message
      outGS.Body.Finished = finished
		}
		err := json.NewEncoder(rw).Encode(gsAnswer)
		return err
	}, requestWaitInQueueTimeout)

	if err != nil {
		//http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}
}
func (session *gameSession) parse(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	defer body.Close()
	err := decoder.Decode(&session)
	fmt.Println(session.Status)
	return err
}

func RunHTTPServer(addr string) error {
	http.HandleFunc("/updateUData", handlers.UpdateUsersData)
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/signin", handlers.SignIn)
	http.HandleFunc("/start", handlers.Start)
	return http.ListenAndServe(addr, nil)
}
