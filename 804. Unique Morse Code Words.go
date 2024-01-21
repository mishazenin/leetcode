package main

import "fmt"

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

var morzeMap = map[byte]string{
	'a': `".-"`,
	'b': `".-"`,
	'c': `".-"`,
	'd': `".-"`,
	'e': `".-"`,
	'f': `".-"`,
	'':  `".-"`,
}

func uniqueMorseRepresentations(words []string) int {

}
