package main

import (
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {

	fw, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fw.WriteString(fmt.Sprint(i, " * ", j, " = ", i*j, "\t"))
		}
		fw.WriteString("\n")
	}

	var g errgroup.Group
	g.Go(test())

	g.Go(func() func() error {
		return func() error {
			return nil
		}
	}())

}

func test() func() error {
	return func() error {
		return nil
	}
}
