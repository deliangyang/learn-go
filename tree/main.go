package main

import (
	"fmt"

	_ "github.com/deliangyang/hello-test/pkg"
)

type Node struct {
	Left  *Node
	Right *Node
	data  string
}

func (n *Node) Insert(data string) {
	if data > n.data {
		if n.Left != nil {
			n.Left.Insert(data)
		} else {
			n.Left = &Node{
				data: data,
			}
		}
	} else {
		if n.Right != nil {
			n.Right.Insert(data)
		} else {
			n.Right = &Node{
				data: data,
			}
		}
	}
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Println("data:", n.data)
	n.Right.InOrder()

}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Println("data:", n.data)

	n.Left.PreOrder()
	n.Right.PreOrder()

}

func (n *Node) EndOrder() {
	if n == nil {
		return
	}

	n.Left.EndOrder()
	n.Right.EndOrder()
	fmt.Println("data:", n.data)
}

type treeNode struct {
	value       string
	left, right *treeNode
}

func main() {

	node := Node{
		data: "A",
	}

	node.Insert("B")
	node.Insert("C")
	node.Insert("D")
	node.Insert("F")
	node.Insert("E")

	node.Insert("G")
	node.Insert("H")
	node.Insert("I")

	n := Node{
		data: "A",
		Left: &Node{
			data: "B",
			Left: &Node{
				data: "D",
			},
			Right: &Node{
				data: "F",
				Left: &Node{
					data: "E",
				},
			},
		},
		Right: &Node{
			data: "C",
			Left: &Node{
				data: "G",
				Left: &Node{
					data: "H",
				},
			},
			Right: &Node{
				data: "I",
			},
		},
	}

	n.InOrder()
	fmt.Println("=======================================")

	node.PreOrder()
	fmt.Println("=======================================")
	node.InOrder()
	fmt.Println("=======================================")
	node.EndOrder()
	fmt.Println("=======================================")
	//n.EndOrder()

	a := 100
	b := 101
	c := 101
	fmt.Printf("%v\n", &a)
	fmt.Printf("%v\n", &b)
	fmt.Printf("%v\n", &c)
	fmt.Printf("%d\n", 15*16+8)

	d := []int{1, 2, 4, 5, 6}
	for i, v := range d {
		fmt.Printf("%d=>%d#%v:%v\n", i, v, &v, &d[i])
	}

}
