// Write a function, maxPathSum, that takes in the root of a binary tree that contains number values. The function should return the maximum sum of any root to leaf path within the tree.

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

func maxPathSum(root *Node) int {
	if root == nil {
		return math.MinInt
	}
	if root.left == nil && root.right == nil {
		return root.value
	}
	return root.value + int(math.Max(float64(maxPathSum(root.left)), float64(maxPathSum(root.right))))
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

	fmt.Println(maxPathSum(a))

}
