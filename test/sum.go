package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 4, 4, 6, 6))
}

func sum(num ...int) int {
	total := 0
	for item := range num {
		total += item
	}
	return total
}