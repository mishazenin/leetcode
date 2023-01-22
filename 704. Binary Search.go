package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2
		if target == nums[middle] {
			return middle
		}
		if target < nums[middle] {
			right = middle - 1
		}
		if target > nums[middle] {
			left = left + 1
		}
	}
	return -1
}
