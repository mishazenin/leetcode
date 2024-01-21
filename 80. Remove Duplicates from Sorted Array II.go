package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeDuplicates2([]int{1, 1, 1, 2, 2, 3}))
}

func removeDuplicatesOpt(nums []int) int {
	writer := 0
	count := 0
	num := nums[0]
	for reader := range nums {
		// 2 cases when we want to write a number to the result
		// 1) New number
		// 2) Count of identical numbers written in the result is less than 2
		if nums[reader] != num || (nums[reader] == num && count < 2) {

			//  if it's a new number
			if nums[reader] != num {
				count = 0
				num = nums[reader]
			}
			nums[writer] = num
			writer++
			count++
		}
	}
	return writer
}

// T O(N^2*nlogN)
func removeDuplicates2(nums []int) int {
	unique := make(map[int]int)
	for idx := range nums {
		num := nums[idx]
		occurrences, found := unique[num]
		if !found {
			unique[num]++
		} else {
			if occurrences < 2 {
				unique[num]++
			}
		}
	}
	nums = nums[:0]
	for num, freq := range unique {
		for i := 0; i < freq; i++ {
			nums = append(nums, num)
		}
	}
	sort.Ints(nums)
	return len(nums)
}
