package Structures

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
