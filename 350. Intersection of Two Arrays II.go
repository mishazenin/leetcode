package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 21}, []int{2, 2}))
}

/*
Так как тут могут быть несколько повторящихся элементов
[1,2,2,1]
[2,2]
то просто мапа с структурами не подходит
поэтому мапа а значени инкрементим и декриментим
*/
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := []int{}
	for _, val := range nums1 {
		m[val]++
	}

	for _, elem := range nums2 {
		if m[elem] > 0 {
			res = append(res, elem)
		}
		m[elem]--
	}
	return res
}
