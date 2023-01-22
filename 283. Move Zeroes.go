package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

// [0,1,0,3,12]
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}
