package main

import "fmt"

// XOR -операция побитового складывания и если элемент повторяется один раз то он останется в маске
func xorSolution(nums []int) int {
	mask := 0
	for _, elem := range nums {
		mask ^= elem
	}
	return mask
}

func main() {
	fmt.Println(xorSolution([]int{2, 2, 1, 1, 3}))
}
