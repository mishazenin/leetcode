package main

import "fmt"

type TreeNode235 struct {
	Val   int
	Left  *TreeNode235
	Right *TreeNode235
}

func lowestCommonAncestor(root, p, q *TreeNode235) *TreeNode235 {

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root

}

func main() {

	fmt.Println(n.Left)
}
