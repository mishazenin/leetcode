package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO not resolved
func main() {
	root := &ListNode{}
	AddNode(root, []int{47, 22, 81, 46, 94, 95, 90, 22, 55, 91, 6, 83, 49, 65, 10, 32, 41, 26, 83, 99, 14, 85, 42, 99, 89, 69, 30, 92, 32, 74, 9, 81, 5, 9})
	pairSum(root)
}
func AddNode(head *ListNode, arr []int) {
	curr := head
	for _, node := range arr {
		curr.Next = &ListNode{Val: node}
		curr = curr.Next
	}

}

func pairSum(head *ListNode) int {
	var maxTwin int
	reversedList := Reverse(head)

	curr1 := head
	curr2 := reversedList

	for curr1 != nil {
		fmt.Println(curr2.Val, curr1.Val)
		maxTwin = max2130(maxTwin, curr2.Val+curr1.Val)
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return maxTwin
}

func Reverse(head *ListNode) *ListNode {

	curr := head
	var prev *ListNode
	var next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func max2130(a, b int) int {
	if a > b {
		return a
	}
	return b
}
