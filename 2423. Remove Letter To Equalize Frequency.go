package main

import "fmt"

func main() {
	fmt.Println(equalFrequency("abcc"))
}

func equalFrequency(word string) bool {
	m := make(map[byte]int)
	for i, j := 0, len(word)-1; i < j; {
		_, found := m[word[i]]
		if found {
			return true
		} else {
			m[word[i]]++
		}
		_, found = m[word[j]]
		if found {
			return true
		} else {
			m[word[j]]++
		}

		i++
		j--
	}
	return true
}
