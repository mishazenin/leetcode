package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{1, 2, 3}))
}

func canJump(nums []int) bool {
	target := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}

	return target == 0
}
