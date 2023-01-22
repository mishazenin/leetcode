package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var res []int

	for _, val := range nums1 {
		m[val] = 1
	}

	for _, val := range nums2 {
		count, found := m[val]
		if !found {
			continue

		}
		if count == 1 {
			res = append(res, val)
		}

		m[val] -= 1

	}
	return res
}
