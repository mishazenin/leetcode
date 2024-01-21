package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
}

func largestGoodInteger(num string) string {
	var result string

	for i := 0; i < len(num)-2; i++ {
		str := num[i : i+3]
		if checkEqual(str) {
			result = max(result, str)
			continue
		}
	}

	return result
}

func checkEqual(str string) bool {
	return str[0] == str[1] && str[1] == str[2]
}

func max(a, b string) string {
	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(b)
	if num1 > num2 {
		return a
	}
	return b
}
