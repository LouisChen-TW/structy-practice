// Write a function, createLinkedList, that takes in an array of values as an argument. The function should create a linked list containing each element of the array as the values of the nodes. The function should return the head of the linked list.

package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

func printLinkedList(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.value)
	printLinkedList(head.next)
}

func createLinkedList(arr []interface{}) *Node {
	head := NewNode(nil)
	cur := head
	for _, v := range arr {
		cur.next = NewNode(v)
		cur = cur.next
	}
	return head.next
}

func main() {
	printLinkedList(createLinkedList([]interface{}{}))
	// h -> e -> y
}
