// Write a function, pairProduct, that takes in an array and a target product as arguments. The function should return an array containing a pair of indices whose elements multiply to the given target. The indices returned must be unique.

// Be sure to return the indices, not the elements themselves.

// There is guaranteed to be one such pair whose product is the target.

package main

import "fmt"

func pairProduct(numbers []int, targetNumber int) []int {
	var store = make(map[int]int)
	for i, v := range numbers {
		complement := targetNumber / v
		if _, ok := store[complement]; ok {
			return []int{i, store[complement]}
		} else {
			store[v] = i
		}
	}
	return []int{}
}

func main() {
	pairProduct([]int{3, 2, 5, 4, 1}, 8)    // -> [1, 3]
	pairProduct([]int{3, 2, 5, 4, 1}, 10)   // -> [1, 2]
	pairProduct([]int{4, 7, 9, 2, 5, 1}, 5) // -> [4, 5]
	fmt.Println(pairProduct([]int{4, 7, 9, 2, 5, 1}, 5))
}
