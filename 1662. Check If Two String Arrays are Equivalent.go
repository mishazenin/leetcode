package main

import "fmt"

// ["ab", "c"]
// ["a", "bc"]
func main() {
	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
}
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var m1 string
	var m2 string

	for _, word := range word1 {
		m1 += word
	}

	for _, word := range word2 {
		m2 += word
	}

	if len(m1) != len(m2) {
		return false
	}
	for i := 0; i < len(m1); i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}
