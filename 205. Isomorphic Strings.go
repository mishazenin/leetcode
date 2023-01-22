package main

import "fmt"

// "egg"
// "add"
func main() {
	isIsomorphic("egg", "add")
}
func isIsomorphic(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		fmt.Println(m1[int(s[i])], m2[int(t[i])])
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}

		i2 := i + 1
		m1[int(s[i])] = i2
		m2[int(t[i])] = i2
	}

	return true
}
