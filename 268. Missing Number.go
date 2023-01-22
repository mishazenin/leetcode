package main

import "fmt"

// сумма арифметической прогрессии
//
//	n * (n+1) /2 - sum(nums)
//
// т.е. от суммы арифм прогресии (сумма элементов как будто там есть все элементы) вычитаем
// текущаю сумму и получаем на выходе недостающее значение
// On по сложности так как ободим весь массив
// 01 - по памяти константа
func missingNumber(nums []int) int {
	n := len(nums)
	return n*(n+1)/2 - sum(nums)
}

func sum(nums []int) int {
	var res int
	for _, i := range nums {
		res += i
	}
	return res
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}
