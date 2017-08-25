package models

import (
	"errors"
	"fmt"
	"io"
	"log"
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
	fmt.Println("correct", correct)
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
		return false, false, "Некорректный ход", p.answer()
	}
	if !p.started {
		return false, false, "Игра не началась", []int{}
	}
	correct, finished, message := p.update(newState)
	return correct, finished, message, p.answer()
}

func (p *Param) update(newState int) (bool, bool, string) { // correct, finished, message
	possibleStates := p.answer()
	correct := false
	for i := 0; i < len(possibleStates) && !correct; i++ {
		correct = possibleStates[i] == newState
	}
	if !correct {
		return false, false, "Некорректный ход"
	}
	p.stepsCount++
	p.curPos = newState - 1
	item := p.gameMap[p.curPos][p.curPos]
	itemFound := false
	message := ""
	if item == 1 {
		itemFound = true
		message = fmt.Sprintf("Сделано ходов: %d, Получен предмет %d", p.stepsCount, item)
		p.items[item-1] = 1
	} else if item > 1 {
		if p.items[item-2] == 1 {
			itemFound = true
			message = fmt.Sprintf("Сделано ходов: %d, Использован предмет %d. Получен предмет %d", p.stepsCount, item-1, item)
			p.items[item-1] = 1
		} else {
			return true, false, fmt.Sprintf("Сделано ходов: %d, Найден предмет %d. Для его получения необходимо использовать предмет %d", p.stepsCount, item, item-1)
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
			return true, false, fmt.Sprintf("Сделано ходов: %d, Найден выход, для открытия которого нужно использовать предмет %d", p.stepsCount, len(p.items))
		}
	}
	return true, false, fmt.Sprintf("Сделано ходов: %d", p.stepsCount)
}
