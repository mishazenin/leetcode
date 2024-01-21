package main

import "fmt"

func main() {
	// [[5],[0,1,2],[3,4,6]]
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}

func groupThePeople(groupSizes []int) [][]int {
	var result [][]int
	groupMapSizes := make(map[int]struct{})
	for _, elem := range groupSizes {
		groupMapSizes[elem] = struct{}{}
	}

	for size := range groupMapSizes {
		group := make([]int, 0, size)
		for idx, guy := range groupSizes {
			if guy == size {
				group = append(group, idx)
			}
			if len(group) == cap(group) {
				result = append(result, group)
				group = make([]int, 0, size)
			}
		}

	}
	return result
}
