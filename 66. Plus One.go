package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}

/*
Проходим по массиву с конца и держим в уме carry при добавлении обнуляем его
*/
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	carry := 1
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i]+carry > 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	if digits[0] == 0 && carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
