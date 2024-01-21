package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{2, 0, 10, 11, 1, 2}
	_ = copy(s, s[2:])
	fmt.Println(s)
	// fmt.Println(countValidWords("cat and  dog"))
	// fmt.Println(strings.Fields("cat and  dog"))
}
func countValidWords(sentence string) int {
	var counter int
	for _, word := range strings.Fields(sentence) {
		if isValidWord(word) {
			counter++
		}
	}
	return counter
}

var punc = map[byte]bool{
	'.': true,
	',': true,
	'!': true,
}

func isValidWord(word string) bool {

	if word[0] == '-' && word[len(word)-1] == '-' {
		return false
	}
	return false
}
