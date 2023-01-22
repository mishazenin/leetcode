package main

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	var res = nums[0]
	if len(nums) > 2 {
		res = nums[2]
	}
	return res
}
