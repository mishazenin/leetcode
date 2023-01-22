package main

import "fmt"

// перебираем числа и проверяем если добавили новое число не стало ли меньше и перезаписываем максимум,
// проверяем сумму всего массива т.к. он является подмассивом самого себя
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	var totalMax = nums[0]
	currSum := nums[0]
	for i := 1; i < len(nums); i++ {
		totalMax += nums[i]
		num := nums[i]
		currSum = maxNum(currSum+num, num)
		maxSum = maxNum(currSum, maxSum)
	}
	fmt.Println(totalMax)
	fmt.Println(maxSum)
	return maxNum(totalMax, maxSum)
}
func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxSubArray([]int{5, 4 - 1, 7, 8})
}
