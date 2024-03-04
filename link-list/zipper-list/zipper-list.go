// Write a function, zipperLists, that takes in the head of two linked lists as arguments. The function should zipper the two lists together into single linked list by alternating nodes. If one of the linked lists is longer than the other, the resulting list should terminate with the remaining nodes. The function should return the head of the zippered linked list.

// Do this in-place, by mutating the original Nodes.

// You may assume that both input lists are non-empty.

package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func zipperLists(head1, head2 *Node) *Node {
	count := 0
	head := head1
	tail := head
	pointer1 := head1.next
	pointer2 := head2
	for pointer1 != nil && pointer2 != nil {
		if count%2 == 0 {
			tail.next = pointer2
			pointer2 = pointer2.next
		} else {
			tail.next = pointer1
			pointer1 = pointer1.next
		}
		tail = tail.next
		count++
	}
	if pointer1 != nil {
		tail.next = pointer1
	}
	if pointer2 != nil {
		tail.next = pointer2
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

	d.next = e
	e.next = f
	fmt.Println(zipperLists(a, d))
}
