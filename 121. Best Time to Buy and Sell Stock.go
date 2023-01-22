package main

import "fmt"

func maxProfit(prices []int) int {
	var maxProfit int
	current_min := prices[0]
	for _, elem := range prices {

		if elem < current_min {
			current_min = elem
		}
		currMax := elem - current_min
		if maxProfit < currMax {
			maxProfit = currMax
		}

	}
	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
