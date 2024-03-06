// Write a function, depthFirstValues, that takes in the root of a binary tree. The function should return an array containing all values of the tree in depth-first order.

// Hey. This is our first binary tree problem, so be extra sure to check out the approach video! -AZ

package main

import "fmt"

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewNode(val string) *Node {
	return &Node{value: val}
}

// func depthFirstValues(root *Node) []string {
// 	if root == nil {
// 		return []string{}
// 	}
// 	var result []string
// 	stack := []*Node{root}
// 	for len(stack) > 0 {
// 		last := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		result = append(result, last.value)

// 		if last.right != nil {
// 			stack = append(stack, last.right)
// 		}
// 		if last.left != nil {
// 			stack = append(stack, last.left)
// 		}
// 	}
// 	return result
// }

func depthFirstValues(root *Node) []string {
	if root == nil {
		return []string{}
	}
	leftValues := depthFirstValues(root.left)
	rightValues := depthFirstValues(root.right)

	result := []string{root.value}
	result = append(result, leftValues...)
	result = append(result, rightValues...)

	return result
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

	fmt.Println(depthFirstValues(a))
}
