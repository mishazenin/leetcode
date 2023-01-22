package main

func twoSumII(numbers []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range numbers {
		i, found := m[target-num]
		if found {
			return []int{i + 1, idx + 1}
		}
		m[num] = idx
	}
	return []int{}
}
