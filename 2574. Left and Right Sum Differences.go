package main

import "fmt"

// [15,1,11,22]
func main() {
	fmt.Println(leftRigthDifference([]int{10, 4, 8, 3}))
}

// T - O(2N) -> O(N)
// S O(N)
func leftRigthDifference(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, 0
	for _, num := range nums {
		right += num
	}
	for i, num := range nums {
		right -= num
		result[i] = Abs(left - right)
		left += num
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
