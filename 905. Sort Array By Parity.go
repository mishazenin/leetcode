package main

// [3,1,2,4]
func sortArrayByParity(nums []int) []int {
	even := []int{}
	odds := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
			continue
		}
		odds = append(odds, num)
	}
	return append(even, odds...)

}
