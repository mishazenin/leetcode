package main

import "fmt"

func main() {
	fmt.Println(differenceOfSum([]int{1, 15, 6, 3}))
}

func differenceOfSum(nums []int) int {
	var elemSum int
	for _, num := range nums {
		elemSum += num - getDigitSum(num)
	}
	return elemSum
}

func getDigitSum(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
