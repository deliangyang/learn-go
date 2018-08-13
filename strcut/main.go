package main

import "fmt"

type human interface {
	getName() string
	getAge() int
}

type person struct {
	age int
	name string
}

func (p person) getName() string {
	return p.name
}

func (p person) getAge() int {
	return p.age
}

func doSomething(h human) {
	fmt.Println("age:", h.getAge())
	fmt.Println("name:", h.getName())
}

func main() {

	a := person{1, "lisi"}
	fmt.Println(a)

	fmt.Println("age:", a.getAge())
	fmt.Println("name:", a.getName())

	doSomething(a)

}
