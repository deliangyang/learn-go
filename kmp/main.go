package main

import (
	"fmt"
	"strings"
)

func getNext(str string) []int {
	next := make([]int, len(str))
	for x := range next {
		next[x] = -1
	}

	fmt.Println(next)

	var i = 0
	var j = -1

	for range str {
		if j == -1 || str[i] == str[j] {
			j += 1
			i += 1
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func kmp(str string, str2 string, next []int) int {
	var i = 0
	var j = 0

	for {
		if i >= len(str) || j >= len(str2) {
			break
		}

		if j == -1 || str[i] == str2[j] {
			i += 1
			j += 1
		} else {
			j = next[j]
		}
	}

	if j == len(str2) {
		return i - j
	}
	return -1
}

func main() {
	var a = "xxxadfsaf2"
	var b = "fsaf"
	strings.Index(a, b)

	next := getNext(a)
	fmt.Println(next)
	fmt.Println(kmp(a, b, next))
	fmt.Println(strings.Index(a, b))

}
