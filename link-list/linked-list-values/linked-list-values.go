// Write a function, linkedListValues, that takes in the head of a linked list as an argument. The function should return an array containing all values of the nodes in the linked list.

package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func fillValues(head *Node, arr *[]string) {
	if head == nil {
		return
	}
	*arr = append(*arr, head.value)
	fillValues(head.next, arr)
}

func linkedListValues(head *Node) []string {
	newArr := []string{}
	fillValues(head, &newArr) // Pass the address of the slice
	return newArr
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	a.next = b
	b.next = c
	c.next = d

	fmt.Println(linkedListValues(a))
}
