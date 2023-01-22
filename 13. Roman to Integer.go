package main

func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	var res int
	for _, letter := range s {
		if num, found := m[letter]; found {
			res += num
		}
	}
	return res

}
