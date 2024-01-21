package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Key][]string)

	for _, v := range strs {
		key := strKey(v)
		groups[key] = append(groups[key], v)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}
