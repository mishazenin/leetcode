package main

import "fmt"

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}

func checkIfPangram(sentence string) bool {
	m := make(map[int32]struct{})
	for _, char := range sentence {
		_, found := m[char]
		if !found {
			m[char] = struct{}{}
		}
	}
	return len(m) == 26
}
