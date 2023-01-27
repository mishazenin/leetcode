package main

func buildArray(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	for idx := 0; idx < len(nums); idx++ {
		res[idx] = nums[nums[idx]]
	}
	return res
}
