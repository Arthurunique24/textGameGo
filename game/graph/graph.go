package graph

import (
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	StartStateFlag = -1
	EndStateFlag   = -2
)

func getMaxItemsCountForStatesCount(statesCount int) int {
	return int(math.Sqrt(float64(statesCount)))
}

func getMaxEdgesForStateWithStatesCount(statesCount int) int {
	return int(math.Sqrt(float64(statesCount)))
}

func placeItems(graph *[][]int, placedStates *[]int, startState int) []int {
	statesCount := len(*graph)
	itemsCount := getMaxItemsCountForStatesCount(statesCount)
	lastState := startState
	log.Printf("Количество генерируемых предметов на карте - %d\n", itemsCount)
	items := make([]int, itemsCount)
	for i := 0; i < itemsCount; i++ {
		lastState = getFarestState(*graph, *placedStates, lastState)
		(*placedStates)[lastState] = 1
		items[i] = lastState
		(*graph)[lastState][lastState] = i + 1
		log.Printf("Генерация %d предмета завершена\n", i+1)
	}
	return items
}

func getFarestState(gameMap [][]int, excl []int, firstState int) int {
	statesCount := len(gameMap)
	q := []int{firstState}
	qStart := 0
	distances := make([]int, statesCount)
	idxMax := q[0]
	maxDistance := 0
	for j := 0; j < statesCount; j++ {
		distances[j] = -1
	}
	distances[q[0]] = 0
	for qStart < len(q) {
		s := q[qStart]
		if maxDistance < distances[s] && excl[s] == 0 {
			maxDistance = distances[s]
			idxMax = s
		}
		qStart++
		for j := 0; j < statesCount; j++ {
			if distances[j] == -1 && gameMap[s][j] == 1 {
				q = append(q, j)
				distances[j] = distances[s] + 1
			}
		}
	}
	log.Printf("Список уже посещённых - %v. Следующее самое удалённое состояние - %d\n", excl, idxMax+1)
	return idxMax
}

func GenerateGraphWithPlacedItems(statesCount int) [][]int {
	rand.Seed(time.Now().Unix())
	graph := make([][]int, statesCount)
	for i := 0; i < statesCount; i++ {
		graph[i] = make([]int, statesCount)
	}
	for i := 1; i < statesCount; i++ {
		log.Printf("Генерация для %d позиции\n", i+1)
		edgesCount := 1 + rand.Intn(getMaxEdgesForStateWithStatesCount(statesCount))
		log.Printf("Позиция %d: Генерация %d рёбер\n", i+1, edgesCount)
		for k := 0; k < edgesCount; k++ {
			log.Printf("Позиция %d: Генерация %d-го рёбра\n", i+1, k+1)
			edgeTo := rand.Intn(i)
			if graph[i][edgeTo] == 0 {
				log.Printf("Позиция %d: Новое ребро %d-%d установлено\n", i+1, i+1, edgeTo+1)
			} else {
				log.Printf("Позиция %d: Ребро %d-%d уже существует\n", i+1, i+1, edgeTo+1)
			}
			graph[i][edgeTo] = 1
			graph[edgeTo][i] = 1
		}
	}
	startState := rand.Intn(statesCount)
	log.Printf("Начальная позиция - %d\n", startState+1)
	placedStates := make([]int, statesCount)
	placedStates[startState] = 1
	graph[startState][startState] = StartStateFlag
	items := placeItems(&graph, &placedStates, startState)
	endState := getFarestState(graph, placedStates, items[len(items)-1])
	log.Printf("Конечная позиция - %d\n", endState+1)
	graph[endState][endState] = EndStateFlag
	return graph
}
