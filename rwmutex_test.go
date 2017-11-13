package rwmutex

import (
	"testing"
)

func TestAddAndShow(t *testing.T) {
	m := NewGoodCounter()
	ch := make(chan int)
	go func(ch chan int) {
		m.AddOne()
		ch <- 1
	}(ch)

	_ = <-ch
	count := m.Show()
	if count != 1 {
		t.Fatalf("bad bad bad count=%d", count)
	}

}

func TestBadCounter(t *testing.T) {
	m := BadCounter{}
	ch := make(chan int)
	go func(ch chan int) {
		m.AddOne()
		ch <- 1
	}(ch)

	_ = <-ch
	count := m.Show()
	if count != 1 {
		t.Fatalf("bad bad bad count=%d", count)
	}

}
