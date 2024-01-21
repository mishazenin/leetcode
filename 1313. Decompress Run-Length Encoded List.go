package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
}

func decompressRLElist(nums []int) []int {
	if (len(nums) % 2) != 0 {
		return nil
	}
	var result []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}
	return result
}
