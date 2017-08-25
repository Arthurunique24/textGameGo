package game

import (
	"fmt"
	"log"

	"../game/graph"
)

type gameSession struct {
	items      []int
	curPos     int
	gameMap    [][]int
	startPos   int
	endPos     int
	stepsCount int
	started    bool
}

const MapSize = 11

var gs gameSession

func Start() []int {
	gs.gameMap = graph.GenerateGraphWithPlacedItems(MapSize)
	log.Println("Сгенерированная карта")
	for i := 0; i < MapSize; i++ {
		for j := 0; j < MapSize; j++ {
			log.Printf("%d ", gs.gameMap[i][j])
		}
		log.Println()
	}

	itemsCount := 0
	for i := 0; i < MapSize; i++ {
		if gs.gameMap[i][i] > itemsCount {
			itemsCount = gs.gameMap[i][i]
		}
	}
	gs.items = make([]int, itemsCount)
	for i := 0; i < MapSize; i++ {
		if gs.gameMap[i][i] == graph.StartStateFlag {
			gs.startPos = i
		} else if gs.gameMap[i][i] == graph.EndStateFlag {
			gs.endPos = i
		} else if gs.gameMap[i][i] > 0 {
			gs.items[gs.gameMap[i][i]-1] = i
		}
	}
	log.Printf("Положение предметов: %v, начальная позиция: %d, конечная позиция; %d\n", gs.items, gs.startPos, gs.endPos)

	gs.stepsCount = 0
	gs.started = true

	return answer()
}

func Turn(newState int) (bool, bool, string, []int) {
	if !gs.started {
		return false, false, "Игра не началась", []int{}
	}
	correct, finished, message := update(newState)
	return correct, finished, message, answer()
}

func update(newState int) (bool, bool, string) { // correct, finished, message
	possibleStates := answer()
	correct := true
	for i := 0; i < len(possibleStates) && correct; i++ {
		correct = possibleStates[i] != newState
	}
	if !correct {
		return false, false, "Некорректный ход"
	}
	gs.stepsCount++
	gs.curPos = newState
	item := gs.gameMap[gs.curPos][gs.curPos]
	itemFound := false
	message := ""
	if item == 1 {
		itemFound = true
		message = fmt.Sprintf("Сделано ходов: %d, Найден предмет %d", gs.stepsCount, item)
		gs.items[item-1] = 1
	} else if item > 1 && gs.items[item-2] == 1 {
		itemFound = true
		message = fmt.Sprintf("Сделано ходов: %d, Использован предмет %d. Найден предмет %d", gs.stepsCount, item-1, item)
		gs.items[item-2] = 1
	}
	if itemFound {
		gs.gameMap[gs.curPos][gs.curPos] = 0
		return true, false, message
	}
	if gs.curPos == gs.endPos && gs.items[len(gs.items)-1] == 1 {
		gs.started = false
		return true, true, fmt.Sprintf("Поздрвляем, вы выбрались за %d ходов", gs.stepsCount)
	}
	return true, false, fmt.Sprintf("Сделано ходов: %d", item)
}

func answer() []int {
	var states []int
	for j := 0; j < MapSize; j++ {
		if gs.gameMap[gs.curPos][j] == 1 {
			states = append(states, j)
		}
	}
	log.Printf("Возможные переходы: %v\n", states)

	return states
}
