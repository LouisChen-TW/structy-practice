// Write a function, getNodeValue, that takes in the head of a linked list and an index. The function should return the value of the linked list at the specified index.

// If there is no node at the given index, then return null.

package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func getNodeValue(head *Node, index int) *string {
	for head != nil {
		if index == 0 {
			return &head.value
		}
		head = head.next
		index--
	}
	return nil
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")

	a.next = b
	b.next = c
	c.next = d

	if result := getNodeValue(a, 2); result != nil {
		fmt.Println(*result)
	} else {
		fmt.Println(result)
	}
}
