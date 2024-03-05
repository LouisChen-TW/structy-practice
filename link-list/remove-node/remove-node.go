// Write a function, removeNode, that takes in the head of a linked list and a target value as arguments. The function should delete the node containing the target value from the linked list and return the head of the resulting linked list. If the target appears multiple times in the linked list, only remove the first instance of the target in the list.

// Do this in-place.

// You may assume that the input list is non-empty.

// You may assume that the input list contains the target.

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

func removeNode(head *Node, targetVal string) *Node {
	if head.value == targetVal {
		return head.next
	}
	found := false
	prev := head
	cur := head.next
	for !found {
		if cur.value == targetVal {
			found = true
			prev.next = cur.next
		} else {
			prev = cur
			cur = cur.next
		}
	}
	return head
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

	printLinkedList(removeNode(a, "c"))
}
