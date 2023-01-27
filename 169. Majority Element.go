package main

func main() {
	majorityElement([]int{3, 2, 3})
}

// more than len(nums)/2 times
func majorityElement(nums []int) int {
	limit := len(nums) / 2
	m := make(map[int]int)
	for _, elem := range nums {
		m[elem]++
	}
	for k, val := range m {
		if val > limit {
			return k
		}
	}
	return 0
}
