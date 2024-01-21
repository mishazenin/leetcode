package main

import "fmt"

func main() {
	fmt.Println(truncateSentence("chopper is not a tanuki", 5) + "|")
}

func truncateSentence(s string, k int) string {
	var (
		counter int
		word    bool
	)
	for i, char := range s {
		if counter == k {
			return s[:i-1]
		}
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			word = true
			continue
		}
		if word {
			counter++
			word = false
		}
	}
	return s
}
