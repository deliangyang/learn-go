package main

import (
	"fmt"
	"sync"
)

func Producer(queue chan<- int)  {
	for i := 0; i < 10; i++ {
		fmt.Println("send:", i)
		queue<- i+10000
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("receive:", v)
	}
}

func main() {
	c := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		Producer(c)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Consumer(c)
		wg.Done()
	}()

	wg.Wait()
}
