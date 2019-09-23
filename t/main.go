package main

import "fmt"

type A struct {
	a int
	age string
}

func main() {
	a := &A{
		a:   0,
		age: "",
	}

	fmt.Println(a)
}
