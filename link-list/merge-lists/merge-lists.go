// Write a function, mergeLists, that takes in the head of two sorted linked lists as arguments. The function should merge the two lists together into single sorted linked list. The function should return the head of the merged linked list.

// Do this in-place, by mutating the original Nodes.

// You may assume that both input lists are non-empty and contain increasing sorted numbers.

package main

import (
	"fmt"
	"math"
)

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

func mergeLists(head1, head2 *Node) *Node {
	newNode := NewNode(math.MinInt)
	tail := newNode
	pointer1 := head1
	pointer2 := head2
	for pointer1 != nil && pointer2 != nil {
		if pointer1.value < pointer2.value {
			tail.next = pointer1
			pointer1 = pointer1.next
		} else {
			tail.next = pointer2
			pointer2 = pointer2.next
		}
		tail = tail.next
	}
	if pointer1 != nil {
		tail.next = pointer1
	}
	if pointer2 != nil {
		tail.next = pointer2
	}
	return newNode.next
}

func main() {
	a := NewNode(5)
	b := NewNode(7)
	c := NewNode(10)
	d := NewNode(12)
	e := NewNode(20)
	f := NewNode(28)

	a.next = b
	b.next = c
	c.next = d
	d.next = e
	e.next = f

	q := NewNode(6)
	r := NewNode(8)
	s := NewNode(9)
	t := NewNode(25)

	q.next = r
	r.next = s
	s.next = t

	printLinkedList(mergeLists(a, q))
}
