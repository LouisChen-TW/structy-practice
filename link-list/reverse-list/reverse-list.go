// Write a function, reverseList, that takes in the head of a linked list as an argument. The function should reverse the order of the nodes in the linked list in-place and return the new head of the reversed linked list.

package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func reverseList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var prev *Node
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")

	a.next = b
	b.next = c
	c.next = d
	d.next = e
	e.next = f

	fmt.Println(*reverseList(a))
}
