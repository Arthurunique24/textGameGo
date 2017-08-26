package models

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"

	"github.com/ChernovAndrey/textGameGo/models/graph"
)

func GameStep(body io.ReadCloser) (Answer, error) {
	req := new(RequestStep)
	Parse(req, body)
	p := params[req.Id]
	if p == nil {
		fmt.Println("GameStep not SignIn")
		return Answer{}, errors.New("not Authorized")
	}
	correct, finished, message, steps := p.turn(req.Step)
	message = fmt.Sprintf("Ты в %d. %s", p.curPos+1, message)
	fmt.Println("correct", correct)
	fmt.Printf("Положение маньяка - %d\n", p.killerPos+1)
	if finished == true {
		delete(params, req.Id)
	} else {
		mu.Lock()
		params[req.Id] = p
		defer mu.Unlock()
	}
	return Answer{Id: req.Id, PossibleSteps: steps, Finished: finished, Message: message}, nil
}

func (p *Param) turn(newState int) (bool, bool, string, []int) {
	log.Printf("Попытка хода в %d", newState)
	if newState < 1 || newState > len(p.gameMap) {
		return false, false, "Некорректный ход", p.answer(p.curPos)
	}
	if !p.started {
		return false, false, "Игра не началась", []int{}
	}
	correct, finished, message := p.update(newState)
	return correct, finished, message, p.answer(p.curPos)
}

func (p *Param) update(newState int) (bool, bool, string) { // correct, finished, message
	possibleStates := p.answer(p.curPos)
	correct := false
	for i := 0; i < len(possibleStates) && !correct; i++ {
		correct = possibleStates[i] == newState
	}
	if !correct {
		return false, false, "Некорректный ход"
	}
	p.stepsCount++
	nextKillerPos := p.randomMove(p.killerPos)
	fmt.Printf("Следующее положение маньяка - %d\n", nextKillerPos+1)
	if p.killerPos == newState-1 && nextKillerPos == p.curPos {
		return true, true, fmt.Sprintf("Сделано ходов: %d, К сожалению, вас нашёл маньяк. Игра окончена", p.stepsCount)
	}
	p.killerPos = nextKillerPos
	p.curPos = newState - 1
	if p.killerPos == p.curPos {
		return true, true, fmt.Sprintf("Сделано ходов: %d, К сожалению, вас нашёл маньяк. Игра окончена", p.stepsCount)
	}
	tip := p.getTipAboutKillerPresence()
	item := p.gameMap[p.curPos][p.curPos]
	itemFound := false
	message := ""
	if item == 1 {
		itemFound = true
		message = fmt.Sprintf("Сделано ходов: %d, Получен предмет %d. %s", p.stepsCount, item, tip)
		p.items[item-1] = 1
	} else if item > 1 {
		if p.items[item-2] == 1 {
			itemFound = true
			message = fmt.Sprintf("Сделано ходов: %d, Использован предмет %d. Получен предмет %d. %s", p.stepsCount, item-1, item, tip)
			p.items[item-1] = 1
		} else {
			return true, false, fmt.Sprintf("Сделано ходов: %d, Найден предмет %d. Для его получения необходимо использовать предмет %d. %s", p.stepsCount, item, item-1, tip)
		}
	}
	if itemFound {
		p.gameMap[p.curPos][p.curPos] = 0
		return true, false, message
	}
	if p.curPos == p.endPos {
		if p.items[len(p.items)-1] == 1 {
			p.started = false
			return true, true, fmt.Sprintf("Поздравляем, вы выбрались за %d ходов", p.stepsCount)
		} else {
			return true, false, fmt.Sprintf("Сделано ходов: %d, Найден выход, для открытия которого нужно использовать предмет %d. %s", p.stepsCount, len(p.items), tip)
		}
	}
	return true, false, fmt.Sprintf("Сделано ходов: %d. %s", p.stepsCount, tip)
}

func (p *Param) getTipAboutKillerPresence() string {
	path := graph.FindNearestPath(p.gameMap, p.curPos, p.killerPos)
	distanceToKiller := len(path)
	tip := "Вокруг царит подозрительная тишина..."
	if p.distanceToKiller >= distanceToKiller {
		if distanceToKiller >= 1 && distanceToKiller <= 2 {
			tip = fmt.Sprintf("Звуки доносятся из %d. Это маньяк! Надо бежать...", path[0]+1)
		} else if distanceToKiller >= 3 && distanceToKiller <= 4 {
			tip = "Что это было? Кажется, маньяк уже близко. Надо торопиться"
		}
	} else {
		tip = "Кажется, маньяк отстал. Но не стоит расслабляться. Нужно продолжать искать выход"
	}
	p.distanceToKiller = distanceToKiller
	return tip
}

func (p *Param) randomMove(fromPos int) int {
	possibleStates := p.answer(fromPos)
	filteredStates := make([]int, len(possibleStates)-1) // Маньяк не может стоять на месте
	j := 0
	for i := 0; i < len(possibleStates); i++ {
		if possibleStates[i] != fromPos+1 {
			filteredStates[j] = possibleStates[i]
			j++
		}
	}
	i := rand.Intn(len(filteredStates))
	return filteredStates[i] - 1
}
