// Write a function, fiveSort, that takes in an array of numbers as an argument. The function should rearrange elements of the array such that all 5s appear at the end. Your function should perform this operation in-place by mutating the original array. The function should return the array.

// Elements that are not 5 can appear in any order in the output, as long as all 5s are at the end of the array.

package main

import "fmt"

func fiveSort(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left] == 5 && nums[right] != 5 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else {
			if nums[left] != 5 {
				left++
			}
			if nums[right] == 5 {
				right--
			}
		}
	}
	return nums
}

func main() {
	fiveSort([]int{12, 5, 1, 5, 12, 7})            // -> [12, 7, 1, 12, 5, 5]
	fiveSort([]int{5, 2, 5, 6, 5, 1, 10, 2, 5, 5}) // -> [2, 2, 10, 6, 1, 5, 5, 5, 5, 5]
	fiveSort([]int{5, 5, 5, 1, 1, 1, 4})           // -> [4, 1, 1, 1, 5, 5, 5]
	fmt.Println(fiveSort([]int{5, 5, 5, 1, 1, 1, 4}))
}
