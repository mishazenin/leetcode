package main

// func getConcatenation(nums []int) []int {
// 	var res = nums
// 	return append(res, nums...)
// }

func getConcatenation(nums []int) []int {
	var res = nums
	for _, num := range nums {
		res = append(res, num)
	}
	return res
}
