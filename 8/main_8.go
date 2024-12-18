package main

import "fmt"

func main() {
	wg := NewSemaphore()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(1)
	}()

	wg.Wait()
	fmt.Println("done")
}

type Semaphore struct {
	count int
	done  chan struct{}
}

func NewSemaphore() *Semaphore {
	return &Semaphore{
		count: 0,
		done:  make(chan struct{}),
	}
}

func (s *Semaphore) Add(count int) {
	s.count += count
}

func (s *Semaphore) Done() {
	s.count--
	if s.count == 0 {
		close(s.done)
	}
}

func (s *Semaphore) Wait() {
	<-s.done
}
