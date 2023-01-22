package main

func mostWordsFound(sentences []string) int {
	var counter int

	for _, sentence := range sentences {
		var wordCounter = 1
		for _, char := range sentence {
			if char == ' ' {
				wordCounter++
			}
		}
		counter = maxWord(counter, wordCounter)

	}
	return counter

}
func maxWord(a, b int) int {
	if a > b {
		return a
	}
	return b
}
