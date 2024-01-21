package main

import "fmt"

func main() {
	fmt.Println(toLowerCase("Hello"))
}

func toLowerCase(s string) string {
	arr := []byte(s)
	for i := range arr {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] + 32
		}
	}
	return string(arr)
}
