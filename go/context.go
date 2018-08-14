package main

import "fmt"

/**
 * golang callback function
 */
type Test struct {
	OnEvicted func(key Key, value interface{})
}

type Key interface{}

func (test *Test) print(key Key, value interface{} ) {
	if test.OnEvicted != nil {
		test.OnEvicted(key, value)
	}
}

func main() {
	test := Test{
		OnEvicted: func(key Key, value interface{}) {
			fmt.Println("hello world")
			fmt.Println(key, value)
		},
	}

	test.print(1, 2)
}
