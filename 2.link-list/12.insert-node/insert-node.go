// Write a function, insertNode, that takes in the head of a linked list, a value, and an index. The function should insert a new node with the value into the list at the specified index. Consider the head of the linked list as index 0. The function should return the head of the resulting linked list.

// Do this in-place.

// You may assume that the input list is non-empty and the index is not greater than the length of the input list.

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

func insertNode(head *Node, value string, index int) *Node {
	newNode := NewNode(value)
	if index <= 0 {
		newNode.next = head
		return newNode
	}
	curIndex := 1
	prev := head
	cur := head.next
	for cur != nil && curIndex <= index {
		if curIndex == index {
			prev.next = newNode
			newNode.next = cur
		} else {
			prev = cur
			cur = cur.next
		}
		curIndex++
	}

	if cur == nil {
		prev.next = newNode
	}

	return head
}

func main() {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")

	a.next = b
	b.next = c
	c.next = d

	printLinkedList(insertNode(a, "x", -1))
}
