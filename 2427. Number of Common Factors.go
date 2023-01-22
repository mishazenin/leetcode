package main

// a = 12, b = 6
func commonFactors(a int, b int) int {
	var counter int
	for i := 0; i < minFactor(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			counter++
		}
	}
	return counter
}
func minFactor(a, b int) int {
	if a < b {
		return a
	}
	return b
}
