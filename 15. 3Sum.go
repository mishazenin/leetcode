package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var result [][]int

	for i, num := range nums {
		if i != len(nums)-1 && num == nums[i+1] {
			continue
		}
		res := findTwoSum(nums[i:], 0-num)
		if res == nil {
			continue
		}
		result = append(result, append([]int{i}, res...))
	}
	return result
}

func findTwoSum(arr []int, target int) []int {
	for i, j := 0, len(arr)-1; i < j; {
		s := arr[i] + arr[j]
		if s == target {
			return []int{i + 1, j + 1}
		}
		if s > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
