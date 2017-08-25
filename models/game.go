package models

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/ChernovAndrey/textGameGo/DAO"
)

const mapSize = 11

var (
	mu     sync.Mutex
	params = make(map[string]*Param)
)

type Param struct {
	Id       string
	StartPos int
	EndPos   int
	KeyPos   int
	HasKey   bool
	CurPos   int
	Matrix   [mapSize][mapSize]int
}

func GameStart(body io.ReadCloser) (AnswerStart, error) {
	session := new(SessionId)
	Parse(session, body)
	if DAO.CheckId(session.Id) == false {
		fmt.Println("ddddd")
		return AnswerStart{}, errors.New("incorrect SessionId")
	}
	p := NewParam()
	mu.Lock()
	params[session.Id] = p
	defer mu.Unlock()
	return AnswerStart{Id: session.Id, PossibleSteps: p.Answer()}, nil
}

func (p *Param) Answer() []int {
	var states []int
	for j := 0; j < mapSize; j++ {
		if p.Matrix[p.CurPos][j] == 1 {
			fmt.Println(j)
			states = append(states, j)
		}
	}
	return states
}

func (p *Param) newMatrix() {
	p.Matrix = [mapSize][mapSize]int{
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	}
}
func NewParam() *Param {
	p := new(Param)
	p.newMatrix()
	p.StartPos = generateRandPosition(mapSize, []int{})
	p.EndPos = generateRandPosition(mapSize, []int{p.StartPos})
	p.KeyPos = generateRandPosition(mapSize, []int{p.StartPos, p.EndPos})
	p.HasKey = false
	p.CurPos = p.StartPos

	fmt.Println("Key:", p.KeyPos, "Start:", p.StartPos, "End:", p.EndPos)

	return p
}

func generateRandPosition(max int, exclusions []int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	placed := false
	pos := rand.Intn(max)

	if len(exclusions) == 0 {
		return pos
	}

	for !placed {
		for j := 0; j < len(exclusions) && !placed; j++ {
			pos = rand.Intn(max)
			placed = pos != exclusions[j]
		}
	}
	return pos
}
