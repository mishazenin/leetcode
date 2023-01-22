package main

func main() {}

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// sliding window
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	var (
		slow = head
		fast = head
	)

	for {
		if fast.Next == nil {
			return slow
		}
		if fast.Next.Next == nil {
			return slow.Next
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

}
