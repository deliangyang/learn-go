package test

import "fmt"

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
