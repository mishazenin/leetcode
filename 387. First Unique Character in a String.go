package main

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, char := range s {
		m[char] += 1
	}
	for idx, elem := range s {
		if m[elem] == 1 {
			return idx
		}
	}
	return -1
}
