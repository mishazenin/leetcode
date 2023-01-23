package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make(map[byte]uint8, len(magazine))
	for _, i := range magazine {
		m[byte(i)]++
	}
	for _, i := range ransomNote {
		val, ok := m[byte(i)]
		if !ok || val < 1 {
			return false
		}
		m[byte(i)]--
	}
	return true
}
