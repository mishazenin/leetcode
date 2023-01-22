package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestStrChain([]string{"b", "ba", "bca", "bda", "bdca"}))
}

func longestStrChain(words []string) int {
	sort.StringsAreSorted(words)
	res := 0
	wordmap := make(map[string]int)
	for _, word := range words {
		if len(word) == 0 {
			wordmap[word] = 1
			continue
		}

		for i, _ := range word {
			cutWord := word[i:]
			cnt, found := wordmap[cutWord]
			if !found {
				continue
			}
			res = maxCnt(len(word), cnt+1)

		}

	}
	return res
}

func maxCnt(a, b int) int {
	if a < b {
		return b
	}

	return a
}
