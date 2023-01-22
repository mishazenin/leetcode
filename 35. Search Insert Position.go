package main

import "fmt"

// not solved
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := (left + right) / 2

	for left <= right {

		middle = (left + right) / 2
		if target == nums[middle] {
			return middle
		}

		if target < middle {
			right = middle
		} else if target > middle {
			left = middle
		}

	}

	return middle
}

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 56, 57, 58, 59}
	fmt.Println(searchInsert(input, 6))
}
