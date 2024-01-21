package main

import "fmt"

func main() {
	fmt.Println(sumOfUnique([]int{10, 6, 9, 6, 9, 6, 8, 7}))
}

func sumOfUnique(nums []int) int {
	var sum int
	m := make(map[int]bool)
	for _, num := range nums {
		_, found := m[num]
		if !found {
			sum += num
			m[num] = false
		} else {
			if sum-num < 0 || m[num] {
				continue
			}
			sum -= num
			m[num] = true
		}
	}
	return sum
}
