package main

// Definition for a binary tree node.
type TreeNode108 struct {
	Val   int
	Left  *TreeNode108
	Right *TreeNode108
}

func main() {
	sortedArrayToBST([]int{1, 2, 3, 4, 5, 56, 6})
}
func sortedArrayToBST(nums []int) *TreeNode108 {

	if len(nums) == 0 {
		return nil
	}

	return ConstructBSTFromArray(nums, 0, len(nums)-1)
}

func ConstructBSTFromArray(nums []int, left, right int) *TreeNode108 {

	if left > right {
		return nil
	}

	midPoint := left - (right-left)/2
	node := &TreeNode108{Val: nums[midPoint]}
	node.Left = ConstructBSTFromArray(nums, left, midPoint-1)
	node.Right = ConstructBSTFromArray(nums, midPoint+1, right)
	return node

}
