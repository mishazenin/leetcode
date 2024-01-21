package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var maxPalidrom string
	var currentPalidrom string
	for i := 0; i < len(s)-1; i++ {
		firstChar := s[i]
		for j := i + 1; j <= len(s)-1; j++ {
			nextChar := s[j]
			if firstChar == nextChar {
				if isPalidrom(s[i : j+1]) {
					currentPalidrom = s[i : j+1]
					if len(currentPalidrom) > len(maxPalidrom) {
						maxPalidrom = currentPalidrom
					}
				}
			}
		}
	}

	if maxPalidrom == "" {
		return s[:1]
	}
	return maxPalidrom
}

func isPalidrom(str string) bool {
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
