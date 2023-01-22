package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

/*
Т.к в условии дано то что элементы в возрастающем порядке
то сравниваем элементы предыдущий с текущим и если разные то инткрементим
и по прошедшему индексу записываем значение
*/
func removeDuplicates(nums []int) int {

	newLengt := 1
	for i := 1; i < len(nums); i++ {

		if nums[i-1] != nums[i] {
			nums[newLengt] = nums[i]
			newLengt++
		}
	}
	return newLengt
}
