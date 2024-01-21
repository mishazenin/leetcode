package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
}

func maximumNumberOfStringPairs(words []string) int {
	var counter int
	m := make(map[string]struct{})
	for _, word := range words {
		opposite := []rune(word)
		opposite[0], opposite[1] = opposite[1], opposite[0]
		_, ok := m[string(opposite)]
		if ok {
			counter++
		} else {
			m[word] = struct{}{}
		}
	}
	return counter
}
