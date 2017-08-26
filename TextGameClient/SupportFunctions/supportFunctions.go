package SupportFunctions

import (
	"TextGameClient/Structures"
	"encoding/json"
	"fmt"
)

func UnmarshalGameSession(jsonStr []byte) (string, []int, bool, string) {
	var gameSession Structures.GameSession

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

func TakeLoginAndPassword(login string, password string) []byte {
	authSession := &Structures.AuthSession{Login: login, Password: password}
	marshaled, err := json.Marshal(authSession)
	if err != nil {
		fmt.Println(err)
	}
	return marshaled
}

func TakeGameStepJson(id string, step int) []byte {
	gameStep := &Structures.GameStep{Id: id, Step: step}
	marshaled, err := json.Marshal(gameStep)
	if err != nil {
		fmt.Println(err)
	}
	return marshaled
}
