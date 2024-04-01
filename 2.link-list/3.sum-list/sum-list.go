// Write a function, sumList, that takes in the head of a linked list containing numbers as an argument. The function should return the total sum of all values in the linked list.

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

// iterative
// func sumList(head *Node) int {
// 	sum := 0
// 	for head != nil {
// 		sum += head.value
// 		head = head.next
// 	}
// 	return sum
// }

// recursive
func sumList(head *Node) int {
	if head == nil {
		return 0
	}
	return head.value + sumList(head.next)
}

func main() {
	x := NewNode(38)
	y := NewNode(4)
	x.next = y

	fmt.Println(sumList(x))
}
