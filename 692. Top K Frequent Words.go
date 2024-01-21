package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}

type word struct {
	times int
	word  string
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for i := range words {
		if _, ok := m[words[i]]; ok {
			m[words[i]]++
		} else {
			m[words[i]] = 1
		}
	}
	var newWords []word
	for w, times := range m {
		newWords = append(newWords, word{
			times: times,
			word:  w,
		})
	}

	sort.Slice(newWords, func(i, j int) bool {
		if newWords[i].times == newWords[j].times {
			return newWords[i].word[0] > newWords[j].word[0]
		}
		return newWords[i].times > newWords[j].times
	})
	var result []string
	for i := range newWords {
		if i < k {
			result = append(result, newWords[i].word)
			continue
		}
		break
	}
	return result

}
