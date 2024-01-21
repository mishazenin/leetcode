package main

import "fmt"

func main() {
	fmt.Println(isPalindrome2(121))
}

func isPalindrome2(x int) bool {
	str := fmt.Sprintf("%d", x)
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
