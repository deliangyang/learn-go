package main

import (
	"fmt"
	"container/list"
)

/**
 *             -- -> --
 */
type Node struct {
	key Key
	next *Node
}

type Stack struct {
	node *Node
	length int
}

type Key interface {}

func NewStack() *Stack {
	s := &Stack{nil, 0}
	return s
}

func (stack *Stack) push(key Key) {
	s := &Node{key: key, next:stack.node}
	stack.length++
	stack.node = s
}

func (stack *Stack) len() int  {
	return stack.length
}

func (stack *Stack) isEmpty() bool {
	return stack.length == 0
}

func (stack *Stack) pop()(interface{}, bool) {
	node := stack.node
	if stack.node == nil {
		return nil, false
	}
	stack.node = stack.node.next
	stack.length--
	return node.key, true
}

func main() {
	stack := NewStack()

	stack.push(1)
	stack.push(2)
	fmt.Println("length:", stack.len())
	key, _ := stack.pop()
	fmt.Println(key)
	fmt.Println(stack.len())
	fmt.Println("is empty:", stack.isEmpty())
	key2, _ := stack.pop()
	fmt.Println(key2)
	fmt.Println("is empty:", stack.isEmpty())
	key1, _ := stack.pop()
	fmt.Println(key1)

	list := list.List{}
	list.PushBack(1)
	list.PushBack(2)
	e := list.Back()
	fmt.Println(e.Value)
	list.Remove(e)
	fmt.Println(list.Back().Value)
	fmt.Println(list.Back().Value)
}
