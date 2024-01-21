package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	// making a map the go way
	m := make(map[byte]int)
	res := 0

	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := m[s[r]]; ok {
			// only update the left pointer if
			// it's behind the last position of the visited character
			l = maxLength(l, index)
		}
		res = maxLength(res, r-l+1)
		m[s[r]] = r + 1
	}
	return res
}

func maxLength(n, m int) int {
	if n > m {
		return n
	}
	return m
}
