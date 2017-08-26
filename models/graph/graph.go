package graph

import (
	"fmt"
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

func getFarestState(graph [][]int, excl []int, firstState int) int {
	statesCount := len(graph)
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
			if distances[j] == -1 && graph[s][j] == 1 {
				q = append(q, j)
				distances[j] = distances[s] + 1
			}
		}
	}
	log.Printf("Список уже посещённых - %v. Следующее самое удалённое состояние - %d\n", excl, idxMax+1)
	return idxMax
}

func CalculateOptimalPath(graph [][]int) []int {
	var startPos, endPos int
	statesCount := len(graph)
	itemsCount := 0
	for i := 0; i < statesCount; i++ {
		if graph[i][i] == StartStateFlag {
			startPos = i
		} else if graph[i][i] == EndStateFlag {
			endPos = i
		} else if graph[i][i] > 0 {
			itemsCount++
		}
	}
	path := []int{startPos}
	order := make([]int, 2+itemsCount)
	orderLength := len(order)
	order[0] = startPos
	order[orderLength-1] = endPos
	for i := 0; i < statesCount; i++ {
		if graph[i][i] > 0 {
			order[graph[i][i]] = i
		}
	}
	fmt.Printf("Порядок обхода: ")
	for i := 0; i < orderLength; i++ {
		fmt.Printf("%d ", order[i]+1)
	}
	fmt.Println()

	for i := 0; i < orderLength-1; i++ {
		path = append(path, FindNearestPath(graph, order[i], order[i+1])...)
	}
	return path
}

func FindNearestPath(graph [][]int, from int, to int) []int {
	if from == to {
		return []int{to}
	}
	statesCount := len(graph)
	q := []int{from}
	qStart := 0
	distances := make([]int, statesCount)
	prev := make([]int, statesCount)
	for j := 0; j < statesCount; j++ {
		distances[j] = -1
	}
	distances[q[0]] = 0
	for qStart < len(q) {
		s := q[qStart]
		qStart++
		for j := 0; j < statesCount; j++ {
			if distances[j] == -1 && graph[s][j] == 1 {
				q = append(q, j)
				distances[j] = distances[s] + 1
				prev[j] = s
			}
		}
	}
	path := make([]int, distances[to])
	i := to
	j := 1
	for j <= distances[to] {
		path[distances[to]-j] = i
		j++
		i = prev[i]
	}
	fmt.Printf("Оптимальный путь между позициями %d и %d - %d ", from+1, to+1, from+1)
	for i := 0; i < distances[to]; i++ {
		fmt.Printf("%d ", path[i]+1)
	}
	fmt.Println()
	return path
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
