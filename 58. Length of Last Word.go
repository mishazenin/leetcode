package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	if len(s) == 1 {
		return 1
	}
	var counter = 1
	var start bool

	for i := len(s) - 1; i >= 0; i-- {

		if Char(s[i]) {
			start = true
			counter++
		}
		if !Char(s[i]) && start {
			return counter
		}

	}
	return counter
}
func IsEmpty(char byte) bool {
	return char == ' '
}

func Char(char byte) bool {
	return ('a' <= char) && (char <= 'z') || ('Z' <= char) && (char <= 'Z')
}
