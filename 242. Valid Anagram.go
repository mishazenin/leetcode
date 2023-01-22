package main

func isAnagram(s string, t string) bool {

	m := make(map[rune]int)

	for _, char := range s {
		m[char] += 1
	}
	for _, char := range t {
		_, found := m[char]
		if !found {
			return false
		}
		m[char] -= 1
		if m[char] < 0 {
			return false
		}
		continue
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}
