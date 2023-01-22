package main

// [-4,-1,0,3,10]
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	start := 0
	end := len(nums) - 1
	index := len(nums) - 1

	for start <= end {
		if abs(nums[start]) >= abs(nums[end]) {
			result[index] = nums[start] * nums[start]
			start++
		} else {
			result[index] = nums[end] * nums[end]
			end--
		}
		index--
	}
	return result
}
