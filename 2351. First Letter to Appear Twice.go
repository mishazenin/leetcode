package main

import "fmt"

func main() {
	fmt.Printf("%v", repeatedCharacter("abccbaacz"))
}
func repeatedCharacter(s string) byte {
	if len(s) == 0 {
		return 0
	}
	m := make(map[int32]int)
	for _, char := range s {
		times, found := m[char]
		if found {
			if times > 0 {
				return byte(char)
			}
			m[char] = times + 1
		} else {
			m[char] = 1
		}
	}
	return 0
}
