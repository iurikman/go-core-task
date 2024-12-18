package main

import (
	"testing"
	"sync"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch3 <- 3
		ch2 <- 2
	}()

	res := mergeChanels(ch1, ch2, ch3)
	exp1 := 3
	exp2 := 2
	a := <-res
	b := <-res

	if a != exp1 {
		t.Errorf("mergeChanels got %d, want %d", res, exp1)
	}

	if b != exp2 {
		t.Errorf("mergeChanels got %d, want %d", res, exp2)
	}
}
