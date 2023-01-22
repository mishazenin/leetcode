package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp

		invertTree(root.Right)
		invertTree(root.Left)
	}
	return root
}
