// Write a function, treeSum, that takes in the root of a binary tree that contains number values. The function should return the total sum of all values in the tree.

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

// iterative
// func treeSum(root *Node) int {
// 	sum := 0
// 	queue := []*Node{root}
// 	for len(queue) > 0 {
// 		first := queue[0]
// 		queue = queue[1:]
// 		sum += first.value
// 		if first.left != nil {
// 			queue = append(queue, first.left)
// 		}
// 		if first.right != nil {
// 			queue = append(queue, first.right)

// 		}
// 	}
// 	return sum
// }

// recursive
func treeSum(root *Node) int {
	if root == nil {
		return 0
	}
	return root.value + treeSum(root.left) + treeSum(root.right)
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

	fmt.Println(treeSum(a))
}
