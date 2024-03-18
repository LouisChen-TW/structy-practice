// Write a function, compress, that takes in a string as an argument. The function should return a compressed version of the string where consecutive occurrences of the same characters are compressed into the number of occurrences followed by the character. Single character occurrences should not be changed.

// 'aaa' compresses to '3a'
// 'cc' compresses to '2c'
// 't' should remain as 't'

package main

import (
	"fmt"
)

func compress(s string) string {
	var result string = ""
	var i int = 0
	var j int = 0
	for j < len(s) {
		if s[i] == s[j] {
			j++
		} else {
			num := j - i
			var addString string
			if num == 1 {
				addString = string(s[i])
			} else {
				addString = fmt.Sprintf("%d%s", num, string(s[i]))
			}
			result += addString
			i = j
		}
	}
	num := j - i
	var addString string
	if num == 1 {
		addString = string(s[i])
	} else {
		addString = fmt.Sprintf("%d%s", num, string(s[i]))
	}
	result += addString
	return result
}

func main() {
	compress("ccaaatsss") // -> '2c3at3s'
	compress("ssssbbz")   // -> "4s2bz"
	compress("ppoppppp")  // -> '2po5p'
	fmt.Println(compress("ppoppppp"))
}
