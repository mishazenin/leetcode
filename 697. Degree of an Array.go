package main

func findShortestSubArray(nums []int) int {
	m := make(map[int]int)
	var maxNum int
	var maxCounter int

	for _, val := range nums {
		m[val] += 1
		if IsGreater(m[val], maxCounter) {
			maxCounter = m[val]
			maxNum = val
		}
	}

	start := false
	finalCounter := 0

	for _, elem := range nums {

		if elem == maxNum {
			maxCounter--
			start = true
		}
		if maxCounter == 0 {
			start = false
		}

		if start {
			finalCounter++
		}
	}
	return finalCounter
}
func IsGreater(a, b int) bool {
	return a > b
}
