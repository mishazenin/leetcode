package main

func runningSum(nums []int) []int {
	res := []int{nums[0]}

	for i := 1; i <= len(nums)-1; i++ {
		num := nums[i] + res[len(res)-1]
		res = append(res, num)
	}
	return res
}
