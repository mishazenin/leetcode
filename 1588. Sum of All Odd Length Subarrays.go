package main

func sumOddLengthSubarrays(arr []int) int {
	var answer int
	n := len(arr)
	for left := 0; left < n; left++ {
		currentSum := 0
		for right := left; right < n; right++ {
			currentSum += arr[right]
			if (right-left+1)%2 == 1 {
				answer += currentSum
			}
		}
	}
	return answer
}
