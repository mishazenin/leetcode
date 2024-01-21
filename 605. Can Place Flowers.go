package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	var zeroCnt int
	for i := range flowerbed {
		if flowerbed[i] == 0 {
			zeroCnt++
		}
	}
	return zeroCnt < n
}
