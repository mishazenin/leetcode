package main

import "fmt"

func main() {
	fmt.Println(reverseBits(00000010100101000001111010011100))
}
func reverseBits(num uint32) uint32 {
	str := fmt.Sprintf("%d", num)
	for i, j := 0, len(str)-1; i < j; {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return 0
}
