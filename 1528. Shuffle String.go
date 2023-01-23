package main

func main() {
	println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}
func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for idx := range indices {
		result[indices[idx]] = s[idx]
	}
	return string(result)
}
