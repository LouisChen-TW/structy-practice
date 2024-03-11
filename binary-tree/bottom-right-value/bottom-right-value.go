// Write a function, bottomRightValue, that takes in the root of a binary tree. The function should return the right-most value in the bottom-most level of the tree.

// You may assume that the input tree is non-empty.

package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{value: val}
}

func bottomRightValue(root *Node) int {
	queue := []*Node{root}
	var current *Node
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
	return current.value
}

func main() {
	a := NewNode(3)
	b := NewNode(11)
	c := NewNode(10)
	d := NewNode(4)
	e := NewNode(-2)
	f := NewNode(1)

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	fmt.Println(bottomRightValue(a))

}
