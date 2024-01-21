package main

import "fmt"

func main() {
	fmt.Println(makeSmallestPalindrome("egcfe"))
}

func makeSmallestPalindrome(s string) string {
	str := []byte(s)
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			if str[i] > str[j] {
				str[i] = str[j]
			} else {
				str[j] = str[i]
			}
		}
		i++
		j--
	}
	return string(str)
}
