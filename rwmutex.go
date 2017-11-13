package rwmutex

import (
	"sync"
)

type GoodCounter struct {
	lock    sync.RWMutex
	counter int
}

func NewGoodCounter() *GoodCounter {
	return &GoodCounter{}
}
func (m *GoodCounter) AddOne() {
	m.lock.Lock()
	m.counter++
	m.lock.Unlock()
}

func (m *GoodCounter) Show() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.counter
}

type BadCounter struct {
	lock    sync.RWMutex
	counter int
}

func (m BadCounter) AddOne() {
	m.lock.Lock()
	m.counter++
	m.lock.Unlock()
}

func (m BadCounter) Show() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.counter
}
