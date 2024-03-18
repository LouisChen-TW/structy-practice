// Write a function, treeValueCount, that takes in the root of a binary tree and a target value. The function should return the number of times that the target occurs in the tree.

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

// recursive
func treeValueCount(root *Node, target int) int {
	if root == nil {
		return 0
	}
	match := 0
	if root.value == target {
		match = 1
	}
	return match + treeValueCount(root.left, target) + treeValueCount(root.right, target)
}

// iterative
// func treeValueCount(root *Node, target int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	stack := []*Node{root}
// 	count := 0

// 	for len(stack) > 0 {
// 		last := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		if last.value == target {
// 			count++
// 		}
// 		if last.right != nil {
// 			stack = append(stack, last.right)
// 		}
// 		if last.left != nil {
// 			stack = append(stack, last.left)
// 		}
// 	}
// 	return count
// }

func main() {
	a := NewNode(12)
	b := NewNode(6)
	c := NewNode(6)
	d := NewNode(4)
	e := NewNode(6)
	f := NewNode(12)

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	fmt.Println(treeValueCount(a, 6))

}
