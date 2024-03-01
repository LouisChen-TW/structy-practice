// Write a function, linkedListFind, that takes in the head of a linked list and a target value. The function should return a boolean indicating whether or not the linked list contains the target.

package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func linkedListFind(head *Node, target string) bool {
	for head != nil {
		if head.value == target {
			return true
		}
		head = head.next
	}
	return false
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")

	a.next = b
	b.next = c
	c.next = d

	fmt.Println(linkedListFind(a, "c"))
}
