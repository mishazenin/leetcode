package main

// "Let's take LeetCode contest"
func reverseWords(s string) string {
	var start bool
	var word string
	var final string
	for _, val := range s {
		if val == ' ' {
			word = ""
			start = false
			continue
		}
		start = true
		word += string(val)

	}

}

func reverseStr(str string) string {
	s := []byte(str)
	for i, j := 0, len(str)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}
