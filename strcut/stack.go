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
}

type Key interface {}

func NewStack() *Stack {
	s := &Stack{nil}
	return s
}

func (stack *Stack) push(key Key) {
	s := &Node{key: key, next:stack.node}
	stack.node = s
}

func (stack *Stack) pop()(interface{}, bool) {
	node := stack.node
	if stack.node == nil {
		return nil, false
	}
	stack.node = stack.node.next
	return node.key, true
}

func main() {
	stack := NewStack()

	stack.push(1)
	stack.push(2)
	key, _ := stack.pop()
	fmt.Println(key)
	key2, _ := stack.pop()
	fmt.Println(key2)
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
