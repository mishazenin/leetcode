package main

import "fmt"

func main() {
	fmt.Println(sumOfMultiples(7))
}

func sumOfMultiples(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			result += i
		}
	}
	return result
}
