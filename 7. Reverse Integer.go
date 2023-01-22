package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	reverse(1534236469)
}

// not solved
func reverse(x int) int {
	fmt.Println(-1 * math.Pow(2, 31))
	fmt.Println(math.Pow(2, 31) - 1)
	if float64(x) < (-1 * math.Pow(2, 31)) {
		return 0
	}
	if float64(x) > math.Pow(2, 31)-1 {
		return 0
	}
	var res string
	if x < 0 {
		x *= -1
		res = "-"
	}

	str := strconv.Itoa(x)
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	i, _ := strconv.Atoi(res)
	fmt.Println(i)
	return i
}
