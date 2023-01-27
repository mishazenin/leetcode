package main

// 8, 1, 2, 2, 3
// brute force
func smallerNumbersThanCurrent(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		var currentNumber int
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] && i != j {
				currentNumber++
			}
		}
		result = append(result, currentNumber)
	}

	return result
}
