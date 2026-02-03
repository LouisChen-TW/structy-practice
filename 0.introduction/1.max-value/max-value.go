// Write a function, maxValue, that takes in array of numbers as an argument. The function should return the largest number in the array.

// Solve this without using any built-in array methods.

// You can assume that the array is non-empty.

package main

import (
	"fmt"
)

func maxValue(nums []float64) float64 {
	maxNumber := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNumber {
			maxNumber = nums[i]
		}
	}
	return maxNumber
}

func main() {
	maxValue([]float64{4, 7, 2, 8, 10, 9}) // -> 10
	maxValue([]float64{10, 5, 40, 40.3})   // -> 40.3
	maxValue([]float64{-5, -2, -1, -11})   // -> -1
	fmt.Println(maxValue([]float64{-5, -2, -1, -11}))
}
