package main

// [2,5,1,3,4,7], n = 3
func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))
	for i, j := 0, n; j < len(nums); {
		res = append(res, nums[i])
		res = append(res, nums[j])
		i++
		j++

	}
	return res
}
