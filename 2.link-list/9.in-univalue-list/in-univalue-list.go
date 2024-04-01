// Write a function, isUnivalueList, that takes in the head of a linked list as an argument. The function should return a boolean indicating whether or not the linked list contains exactly one unique value.

// You may assume that the input list is non-empty.

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func printLinkedList(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	printLinkedList(head.next)
}

// iterative
// func isUniValueList(head *Node) bool {
// 	next := head.next
// 	for next != nil {
// 		if head.value == next.value {
// 			next = next.next
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

// recursive
func isUniValueList(head *Node) bool {
	if head.next == nil {
		return true
	}
	if head.value != head.next.value {
		return false
	}
	return isUniValueList(head.next)
}

func main() {
	a := NewNode(5)
	b := NewNode(5)
	c := NewNode(5)
	d := NewNode(5)
	e := NewNode(5)
	f := NewNode(5)

	a.next = b
	b.next = c
	c.next = d
	d.next = e
	e.next = f

	fmt.Println(isUniValueList(a))
}
