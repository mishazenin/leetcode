package main

import (
	"fmt"
	"sort"
)

// Создание среза xArr включает в себя итерацию по всем точкам и извлечение x-координат. Эта операция занимает O(n), где n - количество точек.
// Сортировка среза xArr с использованием sort.Ints занимает O(n log n) времени, где n - длина среза.
// Итерация по отсортированному xArr и вычисление разницы между соседними x-координатами также занимают O(n).
// O(n log n)

// Срез xArr создается для хранения всех x-координат, что занимает O(n) пространства.
func maxWidthOfVerticalArea(points [][]int) int {
	var maxArea int
	var xArr []int
	for _, arr := range points {
		currX := arr[0]
		xArr = append(xArr, currX)
	}
	sort.Ints(xArr)
	for idx := range xArr {
		if idx+1 > len(xArr)-1 {
			break
		}
		currMaxArr := xArr[idx+1] - xArr[idx]
		if currMaxArr > maxArea {
			maxArea = currMaxArr
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxWidthOfVerticalArea([][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}))
}
