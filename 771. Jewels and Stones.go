package main

func numJewelsInStones(jewels string, stones string) int {
	var total int
	m := make(map[string]int)
	for _, i := range jewels {
		m[string(i)]++
	}
	for _, char := range stones {
		total += m[string(char)]
	}
	return total
}
