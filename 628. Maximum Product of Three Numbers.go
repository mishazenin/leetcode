package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
}

func maximumProduct(nums []int) int {
	res := 1
	if len(nums) <= 3 {
		for _, val := range nums {
			res *= val
		}
		return res
	}

	sort.Ints(nums)
	l := len(nums)
	if nums[l-1] == 0 {
		return 0
	}
	res1 := nums[l-1] * nums[l-2] * nums[l-3]
	res2 := nums[0] * nums[1] * nums[l-1]
	return max(res1, res2)
}
