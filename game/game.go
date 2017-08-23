package game

import (
	"fmt"
	"math/rand"
	"time"
)

type gameInit struct {
	startPos int
	endPos int
	keyPos int
	hasKey bool
	curPos int
	matrix [MAP_SIZE][MAP_SIZE] int
}

const MAP_SIZE = 11
//gofmt

func (gi *gameInit) createMatrix(){
	gi.matrix = [MAP_SIZE][MAP_SIZE]int {
		{0, 1, 0, 0, 0 ,0, 0, 0, 0, 0, 0},
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

var gi gameInit

func Start() (bool, []int){
	fmt.Println("Welcome to quest! \nYour task is to find the key, and get out of the maze. \nGood Luck!\n")
	gi.createMatrix()
	gi.startPos = generateRandPosition(MAP_SIZE, []int {})
	gi.endPos = generateRandPosition(MAP_SIZE, []int {gi.startPos})
	gi.keyPos = generateRandPosition(MAP_SIZE, []int {gi.startPos, gi.endPos})
	gi.hasKey = false
	gi.curPos = gi.startPos

	fmt.Println("Key:", gi.keyPos, "Start:", gi.startPos, "End:", gi.endPos)

	return true, gi.Answer()
}


func (gi *gameInit)Update(newState int) (bool) {
	gi.curPos = newState
	return gi.curPos == gi.endPos
}

func (gi *gameInit)Turn(newState int) (bool, []int) {
	end := gi.Update(newState)
	if end {
		return true, []int{}
	} else {
		return false, gi.Answer()
	}
}


func (gs *gameInit) Answer() ([]int){
	var states []int
	for j := 0; j < MAP_SIZE; j++ {
		if gs.matrix[gs.curPos][j] == 1 {
			fmt.Println(j)
			states = append(states, j)
		}
	}

	return states
}

func generateRandPosition(max int, exclusions []int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	placed := false
	pos := rand.Intn(max)

	if len(exclusions) == 0 {
		return pos
	}

	for !placed {
		for j := 0; j < len(exclusions) && !placed; j++ {
			pos = rand.Intn(max)
			//fmt.Println("alo")
			placed = pos != exclusions[j]
		}
	}
	return pos
}
