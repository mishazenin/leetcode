package main

import "fmt"

func main() {
	fmt.Println(prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
}

func prefixCount(words []string, pref string) int {
	var counter int
	for _, word := range words {
		if hasPrefix(word, pref) {
			counter++
		}
	}
	return counter
}
func hasPrefix(word, prefix string) bool {
	for len(word) < len(prefix) {
		return false
	}
	return word[:len(prefix)] == prefix
}
