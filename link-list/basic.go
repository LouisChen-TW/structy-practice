package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func printLinkedList(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	printLinkedList(head.next)
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	a.next = b
	b.next = c
	c.next = d

	printLinkedList(a)
}
