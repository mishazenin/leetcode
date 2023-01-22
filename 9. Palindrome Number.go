package main

// 121
func isPalindromeNum(x int) bool {
	if x < 0 {
		return false
	}
	arr := []int{}
	for x != 0 {
		arr = append(arr, x%10)
		x /= 10
	}
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
