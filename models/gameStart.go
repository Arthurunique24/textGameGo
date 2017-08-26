package models

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/ChernovAndrey/textGameGo/DAO"
	"github.com/ChernovAndrey/textGameGo/models/graph"
)

const mapSize = 10

var (
	mu     sync.Mutex
	params = make(map[string]*Param)
)

type Param struct {
	items             []int
	curPos            int
	gameMap           [][]int
	startPos          int
	endPos            int
	stepsCount        int
	started           bool
	minimalStepsCount int
	killerPos         int
	distanceToKiller  int
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
		return Answer{Id: session.Id, PossibleSteps: p.answer(p.curPos), Message: "Игра уже началась."}, nil
	}
	p = NewParam()
	mu.Lock()
	params[session.Id] = p
	defer mu.Unlock()
	return Answer{Id: session.Id, PossibleSteps: p.answer(p.curPos), Message: fmt.Sprintf("Игра началась. Попробуй выбраться за %d шагов. Ты в %d", p.minimalStepsCount, p.curPos+1)}, nil
}

func (p *Param) answer(fromPos int) []int {
	var states []int
	for j := 0; j < mapSize; j++ {
		if p.gameMap[fromPos][j] == 1 || fromPos == j {
			states = append(states, j+1)
		}
	}
	log.Printf("Возможные переходы из позиции %d: %v\n", fromPos+1, states)

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
	items := make([]int, itemsCount)
	for i := 0; i < mapSize; i++ {
		if p.gameMap[i][i] == graph.StartStateFlag {
			p.startPos = i
		} else if p.gameMap[i][i] == graph.EndStateFlag {
			p.endPos = i
		} else if p.gameMap[i][i] > 0 {
			items[p.gameMap[i][i]-1] = i + 1
		}
	}
	fmt.Printf("Положение предметов: %v, начальная позиция: %d, конечная позиция; %d\n", items, p.startPos+1, p.endPos+1)

	p.stepsCount = 0
	p.started = true
	p.curPos = p.startPos

	optimalPath := graph.CalculateOptimalPath(p.gameMap)
	p.minimalStepsCount = len(optimalPath) - 1
	fmt.Println("Оптимальный маршрут")
	for i := 0; i < len(optimalPath); i++ {
		fmt.Printf("%d ", optimalPath[i]+1)
	}
	fmt.Println("")

	p.killerPos = items[0] - 1
	p.distanceToKiller = 100 // большое значение, которое заведомо проходит проверку
	fmt.Printf("Положение маньяка - %d\n", p.killerPos+1)

	return p
}
