package main

import (
	"sync"
	"fmt"
)

func main() {
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

	ch := mergeChanels(ch1, ch2, ch3)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func mergeChanels(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(chs))
	for _, c := range chs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
