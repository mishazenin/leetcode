package main

import "sort"

func maximumWealth(accounts [][]int) int {
	var totals []int
	for _, customer := range accounts {
		var wealth int
		for _, money := range customer {
			wealth += money
		}
		totals = append(totals, wealth)
	}
	sort.Ints(totals)
	return totals[len(totals)-1]
}
