package main

// 8ms / 4.1 mb
func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))
	for idx, num := range nums {
		val, found := sumMap[target-num]
		if found {
			return []int{val, idx}
		}
		sumMap[num] = idx
	}
	return nil
}
