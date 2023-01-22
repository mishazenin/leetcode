package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Тут все просто но внимаение возвращаем head
func deleteDuplicates(head *ListNode) *ListNode {

	currNode := head
	for {
		if currNode == nil {
			break
		}
		if currNode.Next == nil {
			break
		}

		if currNode.Val == currNode.Next.Val {
			currNode.Next = currNode.Next.Next
		} else {
			currNode = currNode.Next
		}
	}
	return head
}
