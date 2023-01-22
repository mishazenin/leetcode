package main

import "strconv"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var stack string
	currNode := head
	for currNode != nil {
		stack += strconv.Itoa(currNode.Val)
		currNode = currNode.Next
	}
	i, _ := strconv.ParseInt(stack, 2, 64)

	return int(i)
}
