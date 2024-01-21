package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
}

func largestAltitude(gain []int) int {
	var (
		highest int
		alt     int
	)
	for i := range gain {
		alt += gain[i]
		if alt > highest {
			highest = alt
		}
	}
	return highest
}
