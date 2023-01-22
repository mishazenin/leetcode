package main

// allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
func countConsistentStrings(allowed string, words []string) int {
	m := make(map[rune]struct{})
	for _, char := range allowed {
		m[char] = struct{}{}
	}
	var counter int
	for _, word := range words {
		var charCounter int
		for _, char := range word {
			_, found := m[char]
			if found {
				charCounter++
			}
		}
		if charCounter == len(word) {
			counter++
		}
	}
	return counter
}
