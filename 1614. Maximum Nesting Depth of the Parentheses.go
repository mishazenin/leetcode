package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}
func maxDepth(s string) int {
	str := []byte(s)
	var open = byte('(')
	var closed = byte(')')
	var counter int
	for i, j := 0, len(str)-1; i < j; {
		if str[i] == open {
			counter++
		}
		if str[j] == closed {
			counter++
		}
		i++
		j--
	}
	result := counter / 2
	return result
}
