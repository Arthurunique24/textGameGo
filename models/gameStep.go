package models

import (
	"errors"
	"fmt"
	"io"
)

func GameStep(body io.ReadCloser) (Answer, error) {
	req := new(RequestStep)
	Parse(req, body)
	p := params[req.Id]
	if p == nil {
		fmt.Println("GameStep not SignIn")
		return Answer{}, errors.New("not Authorized")
	}
	correct, finished, _, steps := p.Turn(req.Step)
	fmt.Println("correct", correct)
	if finished == true {
		delete(params, req.Id)
	} else {
		mu.Lock()
		params[req.Id] = p
		defer mu.Unlock()
	}
	return Answer{Id: req.Id, PossibleSteps: steps, Finished: finished}, nil
}

func (p *Param) Turn(newState int) (bool, bool, string, []int) {
	correct, finished, message := p.update(newState)
	return correct, finished, message, p.Answer()
}

func (p *Param) update(newState int) (bool, bool, string) { // correct, finished, message
	possibleStates := p.Answer()
	correct := false
	for i := 0; i < len(possibleStates); i++ {
		//correct = possibleStates[i] != newState
		if possibleStates[i] == newState {
			correct = true
			break
		}
	}
	if !correct {
		return false, false, "Некорректный ход"
	}
	p.StepsCount++
	p.CurPos = newState
	item := p.Matrix[p.CurPos][p.CurPos]
	//	item := gs.gameMap[gs.curPos][gs.curPos]
	//	itemFound := false
	//	message := ""
	/*	if item == 1 {
				itemFound = true
				message = fmt.Sprintf("Сделано ходов: %d, Найден предмет %d", p.StepsCount, item)
		//		gs.items[item-1] = 1
			} else if item > 1 && gs.items[item-2] == 1 {
				itemFound = true
				message = fmt.Sprintf("Сделано ходов: %d, Использован предмет %d. Найден предмет %d", p.StepsCount, item-1, item)
		//		gs.items[item-2] = 1
			}
			if itemFound {
				p.Matrix[p.CurPos][p.CurPos] = 0
				return true, false, message
			}*/
	/*	if p.CurPos == p.EndPos && gs.items[len(gs.items)-1] == 1 {
		//gs.Started = false
		return true, true, fmt.Sprintf("Поздрвляем, вы выбрались за %d ходов", p.StepsCount)
	}*/
	if p.CurPos == p.EndPos && p.HasKey {
		//gs.Started = false
		return true, true, fmt.Sprintf("Поздрвляем, вы выбрались за %d ходов", p.StepsCount)
	}
	if p.CurPos == p.KeyPos {
		p.KeyPos = -1
		p.HasKey = true
		return true, false, fmt.Sprintf("Вы подобрали ключ")
	}
	return true, false, fmt.Sprintf("Сделано ходов: %d", item)
}
