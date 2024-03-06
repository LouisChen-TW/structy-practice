// Write a function, breadthFirstValues, that takes in the root of a binary tree. The function should return an array containing all values of the tree in breadth-first order.

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

// always use iterative way to solve the breadth first b-tree
func breadthFirstValues(root *Node) []string {
	if root == nil {
		return []string{}
	}
	var result []string
	queue := []*Node{root}
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		result = append(result, first.value)

		if first.left != nil {
			queue = append(queue, first.left)
		}
		if first.right != nil {
			queue = append(queue, first.right)
		}
	}
	return result
}

// use container/list to achieve O(1) remove first element
// func breadthFirstValues(root *Node) []string {
// 	if root == nil {
// 		return []string{}
// 	}
// 	var result []string
// 	queue := list.New()
// 	queue.PushBack(root)

// 	for queue.Len() > 0 {
// 		firstElement := queue.Front()
// 		queue.Remove(firstElement)
// 		first := firstElement.Value.(*Node)

// 		result = append(result, first.value)

// 		if first.left != nil {
// 			queue.PushBack(first.left)
// 		}
// 		if first.right != nil {
// 			queue.PushBack(first.right)
// 		}
// 	}
// 	return result
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

	fmt.Println(breadthFirstValues(a))
}
