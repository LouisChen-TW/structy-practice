// Write a function, mostFrequentChar, that takes in a string as an argument. The function should return the most frequent character of the string. If there are ties, return the character that appears earlier in the string.

// You can assume that the input string is non-empty.

package main

import "fmt"

func mostFrequentChar(s string) string {
	hash := make(map[string]int)
	mostFrequentChar := ""

	for i := 0; i < len(s); i++ {
		if _, ok := hash[string(s[i])]; ok {
			hash[string(s[i])]++
		} else {
			hash[string(s[i])] = 1
		}
	}
	for k := range hash {
		if mostFrequentChar == "" || hash[k] > hash[mostFrequentChar] {
			mostFrequentChar = k
		}
	}
	return mostFrequentChar
}

func main() {
	mostFrequentChar("bookeeper")   // -> "e"
	mostFrequentChar("david")       // -> "d"
	mostFrequentChar("abby")        // -> "b"
	mostFrequentChar("mississippi") // -> "i"
	fmt.Println(mostFrequentChar("mississippi"))
}
