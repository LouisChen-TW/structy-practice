// Write a function, pathFinder, that takes in the root of a binary tree and a target value. The function should return an array representing a path to the target value. If the target value is not found in the tree, then return null.

// You may assume that the tree contains unique values.

package main

import (
	"fmt"
	"slices"
	"time"
)

type Node struct {
	value any
	left  *Node
	right *Node
}

func NewNode(val any) *Node {
	return &Node{value: val}
}

func pathFinderHelper(root *Node, target any) *[]any {
	if root == nil {
		return nil
	}
	if root.value == target {
		return &[]any{root.value}
	}
	leftPath := pathFinderHelper(root.left, target)
	rightPath := pathFinderHelper(root.right, target)
	if leftPath != nil {
		*leftPath = append(*leftPath, root.value)
		return leftPath
	}
	if rightPath != nil {
		*rightPath = append(*rightPath, root.value)
		return rightPath
	}
	return nil
}

func pathFinder(root *Node, target any) *[]any {
	result := pathFinderHelper(root, target)
	if result != nil {
		slices.Reverse(*result)
		return result
	} else {
		return nil
	}
}

// not efficient since it create a new array in each call
// func pathFinder(root *Node, target interface{}) []interface{} {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.value == target {
// 		return []interface{}{root.value}
// 	}

// 	leftPath := pathFinder(root.left, target)
// 	rightPath := pathFinder(root.right, target)

// 	if leftPath != nil {
// 		return append([]interface{}{root.value}, leftPath...)
// 	} else if rightPath != nil {
// 		return append([]interface{}{root.value}, rightPath...)
// 	}

//		return nil
//	}
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

	root := NewNode(0)
	curr := root
	for i := 1; i <= 6000; i += 1 {
		curr.right = NewNode(i)
		curr = curr.right
	}
	start := time.Now()
	// fmt.Println(pathFinder(root, 5000))
	pathFinder(root, 5800)
	fmt.Println(time.Since(start))
}
