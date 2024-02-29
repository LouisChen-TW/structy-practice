package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Write a function, uncompress, that takes in a string as an argument. The input string will be formatted into multiple groups according to the following pattern:

// <number><char>

// for example, '2c' or '3a'.
// The function should return an uncompressed version of the string where each 'char' of a group is repeated 'number' times consecutively. You may assume that the input string is well-formed according to the previously mentioned pattern.

// uncompress("2c3a1t"); // -> 'ccaaat'
// uncompress("4s2b"); // -> 'ssssbb'

func uncompress(s string) string {
	var result strings.Builder
	var jPointer int = 0
	var iPointer int = 0
	var stringSlice []string = strings.Split(s, "")
	for range s {
		if _, err := strconv.Atoi(stringSlice[jPointer]); err != nil {
			if number, err := strconv.Atoi(strings.Join((stringSlice[iPointer:jPointer]), "")); err != nil {
				panic("unable to convert string to integer")
			} else {
				result.WriteString(strings.Repeat(stringSlice[jPointer], number))
				jPointer++
				iPointer = jPointer
			}
		} else {
			jPointer++
		}
	}
	return result.String()
}

func main() {
	fmt.Println(uncompress("2c3a1t"))
}
