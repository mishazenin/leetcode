package main

import "fmt"

func isValid(s string) bool {
	// if the string isn't of even length,
	// it can't be valid so we can return early
	if len(s)%2 != 0 {
		return false
	}

	// set up stack and map
	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// loop over string
	for _, r := range s {

		// if the current character is in the open map,
		// put its closer into the stack and continue
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		// otherwise, we're dealing with a closer
		// check to make sure the stack isn't empty
		// and whether the top of the stack matches
		// the current character
		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		// take the last element off the stack
		st = st[:l]
	}

	// if the stack is empty, return true, otherwise false
	return len(st) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
