// Write a function, levelAverages, that takes in the root of a binary tree that contains number values. The function should return an array containing the average value of each level.

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
// func allTreePaths(root *Node) [][]int {
// 	if root == nil {
// 		return [][]int{}
// 	}
// 	var result [][]int
// 	stack := []struct {
// 		node  *Node
// 		level int
// 	}{
// 		{root, 0},
// 	}

// 	for len(stack) > 0 {
// 		cur := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		if len(result) == cur.level {
// 			result = append(result, []int{cur.node.value})
// 		} else {
// 			curLvl := &result[cur.level]
// 			*curLvl = append(*curLvl, cur.node.value)
// 		}

// 		if cur.node.right != nil {
// 			stack = append(stack, struct {
// 				node  *Node
// 				level int
// 			}{cur.node.right, cur.level + 1})
// 		}
// 		if cur.node.left != nil {
// 			stack = append(stack, struct {
// 				node  *Node
// 				level int
// 			}{cur.node.left, cur.level + 1})
// 		}
// 	}
// 	return result
// }

// recursive
func allTreePaths(root *Node, levels *[][]int, levelNum int) {
	if root == nil {
		return
	}

	if len(*levels) == levelNum {
		*levels = append(*levels, []int{root.value})
	} else {
		curLvl := &(*levels)[levelNum]
		*curLvl = append(*curLvl, root.value)
	}

	allTreePaths(root.left, levels, levelNum+1)
	allTreePaths(root.right, levels, levelNum+1)
}

func levelAverages(root *Node) []float64 {
	var levels [][]int
	allTreePaths(root, &levels, 0)
	var levelAverages []float64
	for _, level := range levels {
		total := 0
		for _, item := range level {
			total += item
		}
		average := float64(total) / float64(len(level))
		levelAverages = append(levelAverages, average)
	}
	return levelAverages
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

	fmt.Println(levelAverages(a))
}
