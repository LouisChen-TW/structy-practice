// Write a function, allTreePaths, that takes in the root of a binary tree. The function should return a 2-Dimensional array where each subarray represents a root-to-leaf path in the tree.

// The order within an individual path must start at the root and end at the leaf, but the relative order among paths in the outer array does not matter.

// You may assume that the input tree is non-empty.

package main

import (
	"fmt"
	"time"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewNode(val string) *Node {
	return &Node{value: val}
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

	now := time.Now()
	fmt.Println(allTreePaths(a))
	fmt.Println(time.Since(now))
}

func allTreePaths(root *Node) [][]string {
	if root == nil {
		return [][]string{}
	}
	if root.left == nil && root.right == nil {
		return [][]string{{root.value}}
	}
	left := allTreePaths(root.left)
	right := allTreePaths(root.right)
	var result [][]string
	for _, sub := range left {
		result = append(result, append([]string{root.value}, sub...))
	}
	for _, sub := range right {
		result = append(result, append([]string{root.value}, sub...))
	}
	return result
}
