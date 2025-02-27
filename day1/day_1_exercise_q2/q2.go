package main

import (
	"fmt"
)

type Node struct {
	Value string
	left  *Node
	right *Node
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%s", root.Value)
	PreOrder(root.left)
	PreOrder(root.right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%s", root.Value)
}

func main() {
	a := &Node{Value: "a"}
	b := &Node{Value: "b"}
	c := &Node{Value: "c"}
	minus := &Node{Value: "-", left: b, right: c}
	plus := &Node{Value: "+", left: a, right: minus}

	PreOrder(plus)
	fmt.Printf("\n")
	PostOrder(plus)

}
