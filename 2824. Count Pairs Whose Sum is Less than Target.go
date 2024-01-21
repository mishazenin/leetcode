package main

import (
	"fmt"
	"sort"
)

// O(N^2)
// O(1)
func countPairs(nums []int, target int) int {
	var counter int
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				counter++
			}
		}
	}
	return counter
}

// O(Nlog(N))
func countPairsOpt(nums []int, target int) int {
	sort.Ints(nums)

	ans, left, right := 0, 0, len(nums)-1
	for left < right {
		if nums[right]+nums[left] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}

	return ans
}

func main() {
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))
}
