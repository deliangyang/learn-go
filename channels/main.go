package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan string)

	go func() {
		message <- "ping"
		time.Sleep(time.Second * 2)
		fmt.Println("hello world")
	}()

	msg := <-message
	fmt.Println("message:", msg)
}
