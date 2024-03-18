// Write a function, howHigh, that takes in the root of a binary tree. The function should return a number representing the height of the tree.

// The height of a binary tree is defined as the maximal number of edges from the root node to any leaf node.

// If the tree is empty, return -1.

package main

import (
	"fmt"
	"math"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewNode(val string) *Node {
	return &Node{value: val}
}

func howHigh(node *Node) int {
	if node == nil {
		return -1
	}
	return int(math.Max(float64(howHigh(node.left)), float64(howHigh(node.right)))) + 1
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	fmt.Println(howHigh(a))
}
