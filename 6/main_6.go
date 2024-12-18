package main

import (
	"math/rand"
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		generator(ch)
	}()

	n := <-ch

	wg.Wait()
	fmt.Println(n)
}

func generator(ch chan int) {
	ch <- rand.Intn(10)
}
