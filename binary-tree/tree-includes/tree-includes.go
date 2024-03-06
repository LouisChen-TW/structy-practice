// Write a function, treeIncludes, that takes in the root of a binary tree and a target value. The function should return a boolean indicating whether or not the value is contained in the tree.

package main

import (
	"fmt"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewNode(val string) *Node {
	return &Node{value: val}
}

// iterative
func treeIncludes(root *Node, target string) bool {
	if root == nil {
		return false
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if first.value == target {
			return true
		}
		if first.left != nil {
			queue = append(queue, first.left)
		}
		if first.right != nil {
			queue = append(queue, first.right)
		}
	}
	return false
}

// recursive
// func treeIncludes(root *Node, target string) bool {
// 	if root == nil {
// 		return false
// 	}
// 	if root.value == target {
// 		return true
// 	}
// 	return treeIncludes(root.left, target) || treeIncludes(root.right, target)
// }

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

	fmt.Println(treeIncludes(a, "c"))
}
