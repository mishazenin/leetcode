package main

func main() {
	pivotIndex([]int{1, 7, 3, 6, 5, 6})
}

// [1,7,3,6,5,6]
func pivotIndex(nums []int) int {
	var totalSum int
	leftSum := 0
	for _, num := range nums {
		totalSum += num
	}

	for idx, elem := range nums {
		if leftSum*2+elem == totalSum {
			return idx
		}
		leftSum += elem
	}
	return -1

}
