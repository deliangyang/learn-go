package main

import (
	"fmt"
	"sync"

	"golang.org/x/tour/tree"
)

func main() {
	wg := sync.WaitGroup{}

	message := make(chan int, 10)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		var count = 100
		for i := 0; i < count; i++ {
			message <- i
		}
		close(message)
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for msg := range message {
			fmt.Println("current i:", msg)
		}
		wg.Done()
	}(&wg)

	wg.Wait()
	for i := 0; i < 100; i++ {
		fmt.Printf("%d => %v\n", i, &i)
		i := i
		fmt.Printf("%d => %v\n", i, &i)
	}

	t := tree.New(10)
	fmt.Println(t)
}
