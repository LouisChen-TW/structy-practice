// Write a function, anagrams, that takes in two strings as arguments. The function should return a boolean indicating whether or not the strings are anagrams. Anagrams are strings that contain the same characters, but in any order.

package main

import "fmt"

func anagrams(s1, s2 string) bool {
	var hash = make(map[string]int)

	for i := 0; i < len(s1); i++ {
		char := string(s1[i])
		if _, ok := hash[char]; ok {
			hash[char]++
		} else {
			hash[char] = 1
		}
	}

	for i := 0; i < len(s2); i++ {
		char := string(s2[i])
		if _, ok := hash[char]; ok {
			hash[char]--
		} else {
			return false
		}
	}

	for k := range hash {
		if hash[k] != 0 {
			return false
		}
	}
	return true
}

func main() {
	anagrams("restful", "fluster")           // -> true
	anagrams("cats", "tocs")                 // -> false
	anagrams("monkeyswrite", "newyorktimes") // -> true
	fmt.Println(anagrams("monkeyswrite", "newyorktimes"))
}
