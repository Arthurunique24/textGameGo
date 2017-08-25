package handlers

import (
	"net/http"

	"encoding/json"
	"fmt"
	"time"
	"../models"
	"../server/workers"
)

var wp *workers.Pool = workers.NewPool(5)

func init() {
	wp.Run()
}

const requestWaitInQueueTimeout = time.Second

func Start(rw http.ResponseWriter, req *http.Request) {
	_, err := wp.AddTaskSyncTimed(func() interface{} {
		res, err := models.GameStart(req.Body)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusNotFound)
			return err
		}
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(rw).Encode(res)
		if err != nil {
			fmt.Println("incorrect format response", err)
		}
		return err
	}, requestWaitInQueueTimeout)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Start", err)
	}
}

func Step(rw http.ResponseWriter, req *http.Request) {
	_, err := wp.AddTaskSyncTimed(func() interface{} {
		res, err := models.GameStep(req.Body)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			return err
		}
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(rw).Encode(res)
		if err != nil {
			fmt.Println("incorrect format response", err)
		}
		return err
	}, requestWaitInQueueTimeout)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("Start", err)
	}
}