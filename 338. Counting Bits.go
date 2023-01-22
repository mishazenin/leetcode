package main

func main() {}

// тут идем по числу
func countBits(n int) []int {
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = memo[i>>1] + i%2
	}
	return memo[:n+1]
}
