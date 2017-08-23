package session

import (
	"sync"
)

var (
	mu sync.RWMutex
	m = make(map[int]bool)
)

func Create(userId int){
	mu.Lock()
	defer mu.Unlock()
	m[userId]=true
}

func CheckId(id int) bool{
	mu.RLock()
	defer mu.RUnlock()
	return m[id]
}

func Delete(id int){
	mu.Lock()
	defer mu.Unlock()
	delete(m,id)
}
