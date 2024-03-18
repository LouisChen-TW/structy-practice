// Write a function, treeMinValue, that takes in the root of a binary tree that contains number values. The function should return the minimum value within the tree.

// You may assume that the input tree is non-empty.

package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{value: val}
}

// iterative
// func treeMinValue(root *Node) int {
// 	min := math.MaxInt
// 	stack := []*Node{root}
// 	for len(stack) > 0 {
// 		last := stack[len(stack)-1]
// 		stack = stack[0 : len(stack)-1]
// 		if last.value < min {
// 			min = last.value
// 		}
// 		if last.left != nil {
// 			stack = append(stack, last.left)
// 		}
// 		if last.right != nil {
// 			stack = append(stack, last.right)

// 		}
// 	}
// 	return min
// }

// recursive
func treeMinValue(root *Node) int {
	if root == nil {
		return math.MaxInt
	}
	min := int(math.Min(float64(treeMinValue(root.left)), float64(treeMinValue(root.right))))
	return int(math.Min(float64(root.value), float64(min)))
}

func main() {
	a := NewNode(3)
	b := NewNode(11)
	c := NewNode(4)
	d := NewNode(4)
	e := NewNode(-2)
	f := NewNode(1)

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	fmt.Println(treeMinValue(a))

}
