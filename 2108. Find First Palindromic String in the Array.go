package main

import "fmt"

func main() {
	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
}

func firstPalindrome(words []string) string {
	for i := range words {
		if isPalidrom(words[i]) {
			return words[i]
		}
	}
	return ""
}

// func isPalidrom(str string) bool {
// 	for i, j := 0, len(str)-1; i < j; {
// 		if str[i] != str[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }
