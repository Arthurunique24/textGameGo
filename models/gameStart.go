package models

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sync"

	"../DAO"
	"../game/graph"
)

const mapSize = 11

var (
	mu     sync.Mutex
	params = make(map[string]*Param)
)

type Param struct {
	items      []int
	curPos     int
	gameMap    [][]int
	startPos   int
	endPos     int
	stepsCount int
	started    bool
}

func GameStart(body io.ReadCloser) (Answer, error) {
	session := new(SessionId)
	Parse(session, body)
	if DAO.CheckId(session.Id) == false {
		fmt.Println("ddddd")
		return Answer{}, errors.New("incorrect SessionId")
	}
	p := params[session.Id]
	if p != nil && p.started {
		fmt.Println("game already started")
		return Answer{Id: session.Id, PossibleSteps: p.answer(), Message: "Игра уже началась"}, nil
	}
	p = NewParam()
	mu.Lock()
	params[session.Id] = p
	defer mu.Unlock()
	return Answer{Id: session.Id, PossibleSteps: p.answer(), Message: "Игра началась"}, nil
}

func (p *Param) answer() []int {
	var states []int
	for j := 0; j < mapSize; j++ {
		if p.gameMap[p.curPos][j] == 1 && p.curPos != j {
			states = append(states, j+1)
		}
	}
	log.Printf("Возможные переходы: %v\n", states)

	return states
}

func NewParam() *Param {
	p := new(Param)
	p.gameMap = graph.GenerateGraphWithPlacedItems(mapSize)
	log.Println("Сгенерированная карта")
	for i := 0; i < mapSize; i++ {
		for j := 0; j < mapSize; j++ {
			fmt.Printf("%2d ", p.gameMap[i][j])
		}
		fmt.Println()
	}

	itemsCount := 0
	for i := 0; i < mapSize; i++ {
		if p.gameMap[i][i] > itemsCount {
			itemsCount = p.gameMap[i][i]
		}
	}
	p.items = make([]int, itemsCount)
	for i := 0; i < mapSize; i++ {
		if p.gameMap[i][i] == graph.StartStateFlag {
			p.startPos = i
		} else if p.gameMap[i][i] == graph.EndStateFlag {
			p.endPos = i
		} else if p.gameMap[i][i] > 0 {
			p.items[p.gameMap[i][i]-1] = i + 1
		}
	}
	fmt.Printf("Положение предметов: %v, начальная позиция: %d, конечная позиция; %d\n", p.items, p.startPos+1, p.endPos+1)

	p.stepsCount = 0
	p.started = true
	p.curPos = p.startPos

	return p
}
