package main

import (
	"sort"
)

func main() {
	_ = longestCommonPrefix([]string{"ab", "a"})
}

func longestCommonPrefix(strs []string) string {
	var longestprefix = ""
	var endPrefix bool

	if len(strs) > 0 {

		sort.Strings(strs)
		first := strs[0]
		last := strs[len(strs)-1]

		for i := 0; i < len(first); i++ {
			if !endPrefix && first[i] == last[i] {
				longestprefix += string(first[i])
			} else {
				endPrefix = true
			}

		}
	}
	return longestprefix
}
