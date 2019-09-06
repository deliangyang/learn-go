package main

import (
	"fmt"
	"sync"
)

type T interface {
	t()
}

type Test interface {
	print()
	T
}

type demo struct {
}

func (d demo) print() {
	fmt.Println(2)
}

func (d demo) t() {
	fmt.Println(3)
}

type Person struct {
	age int
}

func (p Person) Len() {
	return
}

type newInt []int

func (a newInt) Len() int {
	return len(a)
}

func (a newInt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a newInt) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {

	var wg = sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i: ", i)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("=====================")

}

func increment(i int) {
	i++
}

func incrementByPointer(i *int) {
	*i++
}

func printAnyThing(a interface{}) {
	fmt.Println(a)
	switch v := a.(type) {
	case *string:
		fmt.Println(a, ": is string")
		break
	case *int:
		fmt.Println(a, ": is integer")
		break
	default:
		fmt.Println("find nothing: ", v)
	}
}
