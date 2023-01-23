package main

import "fmt"

func subtractProductAndSum(n int) int {
	multiplyRes := 1
	var sumRes int
	nums := parseInt(n)
	for _, digit := range nums {
		multiplyRes *= digit
		sumRes += digit
	}
	fmt.Println(nums, multiplyRes, sumRes)
	return multiplyRes - sumRes

}
func parseInt(input int) []int {
	var res []int
	for input > 0 {
		res = append(res, input%10)
		input = input / 10
	}
	return res
}
