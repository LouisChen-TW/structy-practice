// Write a function, intersection, that takes in two arrays, a,b, as arguments. The function should return a new array containing elements that are in both of the two arrays.

// You may assume that each input array does not contain duplicate elements.

package main

import "fmt"

func intersection(a, b []int) []int {
	newSet := make(map[int]struct{})
	intersection := []int{}
	for _, v := range a {
		newSet[v] = struct{}{}
	}
	for _, v := range b {
		if _, has := newSet[v]; has {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func main() {
	intersection([]int{4, 2, 1, 6}, []int{3, 6, 9, 2, 10}) // -> [2,6]
	intersection([]int{2, 4, 6}, []int{4, 2})              // -> [2,4]
	intersection([]int{4, 2, 1}, []int{1, 2, 4, 6})        // -> [1,2,4]
	fmt.Println(intersection([]int{4, 2, 1}, []int{1, 2, 4, 6}))
}
