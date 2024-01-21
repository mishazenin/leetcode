package main

import "fmt"

func main() {
	fmt.Println(findWordsContaining([]string{"leet", "code"}, 'e'))
}

// Временная сложность функции составляет O(n * m), потому что в худшем случае может потребоваться проверка каждого символа в каждом слове
// В худшем случае, если каждое слово содержит символ x, размер среза res будет равен количеству слов во входном срезе words. Следовательно, пространственная сложность составляет O(n).
func findWordsContaining(words []string, x byte) []int {
	var res []int
	for idx, word := range words {
		for _, char := range word {
			if byte(char) == x {
				res = append(res, idx)
				break
			}
		}
	}
	return res
}
