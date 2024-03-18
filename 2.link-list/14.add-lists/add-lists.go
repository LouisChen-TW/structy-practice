// Write a function, addLists, that takes in the head of two linked lists, each representing a number. The nodes of the linked lists contain digits as values. The nodes in the input lists are reversed; this means that the least significant digit of the number is the head. The function should return the head of a new linked listed representing the sum of the input lists. The output list should have its digits reversed as well.

// You must do this by traversing through the linked lists once.

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

func addLists(head1, head2 *Node) *Node {
	pointer1 := head1
	pointer2 := head2
	head := NewNode(0)
	tail := head
	carry := 0

	for pointer1 != nil || pointer2 != nil {
		pointer1Val := 0
		pointer2Val := 0
		if pointer1 != nil {
			pointer1Val = pointer1.value
			pointer1 = pointer1.next
		}
		if pointer2 != nil {
			pointer2Val = pointer2.value
			pointer2 = pointer2.next
		}
		remain := carry
		sum := pointer1Val + pointer2Val + remain
		carry = 0
		if sum > 9 {
			tail.next = NewNode(sum - 10)
			carry++
		} else {
			tail.next = NewNode(sum)
		}
		tail = tail.next
	}
	if carry > 0 {
		tail.next = NewNode(carry)
	}
	return head.next
}

func main() {
	a := NewNode(1)
	b := NewNode(4)
	c := NewNode(5)
	d := NewNode(7)
	a.next = b
	b.next = c
	c.next = d

	x := NewNode(2)
	y := NewNode(3)
	x.next = y

	a1 := NewNode(9)
	b1 := NewNode(9)
	c1 := NewNode(9)
	a1.next = b1
	b1.next = c1

	x1 := NewNode(6)

	printLinkedList(addLists(a1, x1))
}
