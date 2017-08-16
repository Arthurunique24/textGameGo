package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAP_SIZE = 11

func createMatrix() [MAP_SIZE][MAP_SIZE]int{
	return [11][11]int {
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

func generateRandPosition(max int, exlusions []int) int{
	placed := false
	var pos int = rand.Intn(max)
	if len(exlusions) == 0 {
		return pos
	}
	for !placed {
		for j := 0; j < len(exlusions) && !placed; j++ {
			pos = rand.Intn(max)
			fmt.Println("alo")
			placed = pos != exlusions[j]

		}
	}
	return pos
}

func main() {
	hasKey:=false
	matrix :=createMatrix()
	fmt.Println(matrix)
	fmt.Println("start")
	rand.Seed(time.Now().UTC().UnixNano())
	start := generateRandPosition(MAP_SIZE, []int {})
	end := generateRandPosition(MAP_SIZE, []int {start})
	keyPos := generateRandPosition(MAP_SIZE, []int {start, end})
	fmt.Println("start conditional:","key:",keyPos,"start:",start,"end:",end)
	cur := start
	for !(cur == end && hasKey) {
		var states []int
		for j := 0; j < MAP_SIZE; j++ {
			if matrix[cur][j] == 1 {
				fmt.Println(j)
				states = append(states, j)
			}
		}
		var in int
		ok := false
		for !ok {
			fmt.Scanf("%d", &in)
			for j := 0; j < len(states) && !ok; j++ {
				ok = states[j] == in
			}
		}
		cur = in
		if cur == keyPos {
			hasKey = true
			fmt.Println("found")
		}
	}
}
