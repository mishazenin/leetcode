package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode

	curr := head

	for curr != nil {

		if curr.Val == val {
			if prev == nil {
				head = curr.Next
			} else {
				prev.Next = curr.Next
			}
		}
		prev = curr

		curr = curr.Next

	}

	return head
}
