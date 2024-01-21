package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	var counter int
	input := []rune(s)
	var (
		leftCnt  int
		rightCnt int
	)
	for i := 0; i < len(input); i++ {
		if input[i] == 'L' {
			leftCnt++
		} else {
			rightCnt++
		}
		if leftCnt == rightCnt {
			counter++
			leftCnt = 0
			rightCnt = 0
		}

	}
	return counter
}
