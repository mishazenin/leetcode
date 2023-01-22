package main

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

// Reverse an entire array
// Reverse fisrt half
// Reverse second half
func rotate(nums []int, k int) {
	if k < 1 || len(nums) == 0 || k == len(nums) {
		return
	}
	if k > len(nums) {
		/* We are seeting this instead of k - len(nums)
		   // to handle out if bound
		   Example - nums := [1,2]
		           k = 5
		   if we k - len(nums) it'll be 5 - 2 = 3
		   but length of our input is 2
		*/
		k = k % len(nums)
	}
	reverseArray(nums, 0, len(nums)-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, len(nums)-1)
}

func reverseArray(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
