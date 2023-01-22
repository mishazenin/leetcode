package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

// Помним что 2^8 = 256 и проходимся по всем байтам и проверяем counter
func findTheDifference(s string, t string) byte {
	shortMap := make(map[byte]int)
	longMap := make(map[byte]int)

	for _, val := range s {
		shortMap[byte(val)] += 1
	}

	for _, val := range t {
		longMap[byte(val)] += 1
	}

	for i := 0; i < 256; i++ {
		if longMap[byte(i)] != shortMap[byte(i)] {
			return byte(i)
		}
	}
	return 0
}
