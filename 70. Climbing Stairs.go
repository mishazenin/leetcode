package main

// смысл в том что бы найти решение между n1 и n2
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	n1 := 1
	n2 := 2
	for i := 3; i < n+1; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}

func main() {

}
