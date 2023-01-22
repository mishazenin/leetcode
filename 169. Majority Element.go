package main

import (
	"fmt"
)

func main() {
	majorityElement([]int{3, 2, 3})
}

// more than len(nums)/2 times
func majorityElement(nums []int) int {
	var counter = len(float64(nums)) / 2
	fmt.Println(len(nums) / 2)
	m := make(map[float64]float64)

	for _, num := range nums {
		m[float64(num)]++
	}

	// for cnt, val := range m {
	// 	if cnt >= counter {
	// 		return int(val)
	// 	}
	// }
	return 0

}
