package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}

type wrap struct {
	word   string
	number int
}

func sortSentence(s string) string {
	wrapArr := []wrap{}
	arr := strings.Fields(s)
	for i := range arr {
		wrapArr = append(wrapArr, getIdxFromString(arr[i]))
	}
	sort.Slice(wrapArr, func(i, j int) bool {
		return wrapArr[i].number < wrapArr[j].number
	})

	var result strings.Builder
	for i := range wrapArr {
		result.WriteString(wrapArr[i].word)
		if i != len(wrapArr)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}

func getIdxFromString(str string) wrap {
	idx, _ := strconv.Atoi(string(str[len(str)-1]))
	return wrap{
		word:   str[:len(str)-1],
		number: idx,
	}
}
