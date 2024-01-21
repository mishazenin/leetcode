package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}
func findDifference(nums1 []int, nums2 []int) [][]int {
	var arr1 []int
	var arr2 []int

	uniqueNums2 := make(map[int]struct{})
	for _, e := range nums2 {
		uniqueNums2[e] = struct{}{}
	}
	localMap2 := make(map[int]struct{})
	for _, elem := range nums1 {
		_, found := uniqueNums2[elem]
		if !found {
			if _, ok := localMap2[elem]; !ok {
				arr1 = append(arr1, elem)
				localMap2[elem] = struct{}{}
			}
		}
	}

	uniqueNums1 := make(map[int]struct{})
	for _, e := range nums1 {
		uniqueNums1[e] = struct{}{}
	}
	localMap1 := make(map[int]struct{})
	for _, elem := range nums2 {
		_, found := uniqueNums1[elem]
		if !found {
			if _, ok := localMap1[elem]; !ok {
				arr2 = append(arr2, elem)
				localMap1[elem] = struct{}{}
			}
		}
	}

	return [][]int{arr1, arr2}
}
