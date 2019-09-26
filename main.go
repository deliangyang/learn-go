package main

import (
	"fmt"
	"os"
)

type f func () string

func test (f2 f) {
	fmt.Println(f2())
}

func main() {
	sign := make(chan os.Signal, 1)

	go func() {
		<- sign
	}()

	test := []struct {
		name int
	}{{2}}
}
