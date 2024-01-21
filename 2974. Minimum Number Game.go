package main

import (
	"fmt"
	"sort"
)

// T - O(NlogN + logN) но тк NlogN -> NlogN
// Space O(N)
func numberGame(nums []int) []int {
	var result []int
	sort.Ints(nums)
	for idx := 0; idx < len(nums)-1; idx += 2 {
		result = append(result, nums[idx+1])
		result = append(result, nums[idx])
	}
	return result
}

func main() {
	fmt.Println(numberGame([]int{5, 4, 2, 3}))
}
