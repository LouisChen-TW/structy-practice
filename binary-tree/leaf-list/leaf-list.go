// Write a function, leafList, that takes in the root of a binary tree and returns an array containing the values of all leaf nodes in left-to-right order.

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
	fmt.Println(leafList(a))
	fmt.Println(time.Since(now))
}

// iterative
// func leafList(root *Node) []string {
// 	if root == nil {
// 		return []string{}
// 	}
// 	var list []string
// 	stack := []*Node{root}

// 	for len(stack) > 0 {
// 		cur := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		if cur.left == nil && cur.right == nil {
// 			list = append(list, cur.value)
// 		}
// 		if cur.right != nil {
// 			stack = append(stack, cur.right)
// 		}
// 		if cur.left != nil {
// 			stack = append(stack, cur.left)
// 		}
// 	}
// 	return list
// }

// recursive
func leafList(root *Node) *[]string {
	if root == nil {
		return &[]string{}
	}
	var result []string
	left := leafList(root.left)
	right := leafList(root.right)
	result = append(result, *left...)
	result = append(result, *right...)
	if root.left == nil && root.right == nil {
		result = append(result, root.value)
		return &result
	}
	return &result
}
