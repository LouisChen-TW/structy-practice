// Write a function, pairSum, that takes in an array and a target sum as arguments. The function should return an array containing a pair of indices whose elements sum to the given target. The indices returned must be unique.

// Be sure to return the indices, not the elements themselves.

// There is guaranteed to be one such pair that sums to the target.

package main

import (
	"fmt"
)

func pairSum(numbers []int, targetSum int) []int {
	var store = make(map[int]int)
	for i, v := range numbers {
		fmt.Println(v)
		complement := targetSum - v
		if _, ok := store[complement]; ok {
			return []int{i, store[complement]}
		} else {
			store[numbers[i]] = i
		}
	}
	return []int{}
}

func main() {
	pairSum([]int{3, 2, 5, 4, 1}, 8)    // -> [0, 2]
	pairSum([]int{4, 7, 9, 2, 5, 1}, 5) // -> [0, 5]
	pairSum([]int{4, 7, 9, 2, 5, 1}, 3) // -> [3, 5]
	fmt.Println(pairSum([]int{3, 2, 5, 4, 1}, 8))
}
