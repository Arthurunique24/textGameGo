package models

type SessionId struct {
	Id string
}

type gameSession struct {
	Status string `json:"status,omitempty"`
	Body   struct {
		Turn   int   `json:"turn,omitempty"`
		Answer []int `json:"answer,omitempty"`
	} `json:"body,omitempty"`
}

type Answer struct {
	Id            string
	PossibleSteps []int
	Finished      bool
	Message       string
}

type RequestStep struct {
	Id   string
	Step int
}
