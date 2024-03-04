// Write a function, longestStreak, that takes in the head of a linked list as an argument. The function should return the length of the longest consecutive streak of the same value within the list.
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

func longestStreak(head *Node) int {
	if head == nil {
		return 0
	}
	maxStreak := 1
	curStreak := 1
	curValue := head.value
	pointer := head.next
	for pointer != nil {
		if curValue == pointer.value {
			curStreak++
		} else {
			curStreak = 1
			curValue = pointer.value
		}
		if curStreak > maxStreak {
			maxStreak = curStreak
		}
		pointer = pointer.next
	}

	return maxStreak
}

func main() {
	a := NewNode(5)
	b := NewNode(5)
	c := NewNode(7)
	d := NewNode(7)
	e := NewNode(7)
	f := NewNode(6)

	a.next = b
	b.next = c
	c.next = d
	d.next = e
	e.next = f

	x := NewNode(5)

	fmt.Println(longestStreak(a))
	fmt.Println(longestStreak(x))
	fmt.Println(longestStreak(nil))
}
