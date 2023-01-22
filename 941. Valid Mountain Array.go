package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}

// [2,0,2]

func validMountainArray(arr []int) bool {

	if len(arr) == 1 {
		return false
	}
	if len(arr) == 2 {
		return arr[0] < arr[1]
	}

	for i, j := 0, len(arr)-1; i < j; {

		if (arr[i] >= arr[i+1] && arr[j] >= arr[j-1]) || i+1 == len(arr)-1 || j-1 == 0 {
			return false
		}
		if arr[i] < arr[i+1] {
			i++
		}

		if arr[j] < arr[j-1] {
			j--
		}
	}

	return true
}
