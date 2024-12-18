package main

import (
	"sync"
	"fmt"
	"math"
)

func main() {
	ch1 := make(chan uint8, 10)
	ch2 := make(chan float64, 10)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("sending", i)
			ch1 <- uint8(i)
		}
		close(ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		convertAndTransfer(ch1, ch2)
	}()

	wg.Wait()
	close(ch2)

	for v := range ch2 {
		fmt.Println("reading:", v)
	}
}

func convertAndTransfer(ch1 chan uint8, ch2 chan float64) {
	for a := range ch1 {
		ch2 <- math.Pow(float64(a), 3)
	}
}
