package main

import (
	"fmt"
)

func main() {
	fmt.Println(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
}

func deleteGreatestValue(grid [][]int) int {
	if len(grid) == 1 {
		var max int
		for i := range grid[0] {
			if grid[0][i] > max {
				max = grid[0][i]
			}
		}
		return max
	}
	var counter int
	for i := range grid {
		grid[i] = bubbleSort(grid[i])
	}

	var max int
	for {
		for i, row := range grid {
			newMax, newRow := getMaxFromRow(row)
			if newMax > max {
				max = newMax
			}
			grid[i] = newRow
		}
		counter += max
		max = 0
		if len(grid[0]) == 1 {
			if grid[0][0] > grid[1][0] {
				counter += grid[0][0]
			} else {
				counter += grid[1][0]
			}

			break
		}
	}
	return counter
}

func getMaxFromRow(arr []int) (int, []int) {
	max := arr[len(arr)-1]
	if len(arr) > 1 {
		arr = arr[:len(arr)-1]
	}
	return max, arr
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
