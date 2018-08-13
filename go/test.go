package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("hi start")

	a := 1
	go func(a *int) {
		*a = 2
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}(&a)

	time.Sleep(time.Second)
	fmt.Println("a:", a)
	fmt.Println("hi")
}
