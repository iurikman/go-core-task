package main

import "testing"

func TestGenerator(t *testing.T) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		generator(ch)
	}()
	n := -1
	n = <-ch
	if n == -1 {
		t.Errorf("generator didn't produce output number")
	}
}
