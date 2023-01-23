package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}
func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for idx := range nums {
		movedIndx := index[idx]
		if idx != movedIndx {
			for mdx := idx; mdx > movedIndx; mdx-- {
				target[mdx] = target[mdx-1]
			}
		}
		target[movedIndx] = nums[idx]
	}
	return target
}

// func createTargetArray(nums []int, index []int) []int {
// 	result := make([]int, len(nums), len(nums))
// 	end := make([]int, 0, 0)
// 	for i, val := range index {
// 		v := result[val]
// 		if v == 0 {
// 			result[val] = nums[i]
// 		} else {
// 			end = result[val:]
// 			result = append(result[:val], nums[i])
// 			result = append(result, end...)
// 		}
// 	}
// 	return result
// }
