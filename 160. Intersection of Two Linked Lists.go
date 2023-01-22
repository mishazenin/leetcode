package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	node := headA
	for node != nil {
		node = node.Next
		m[node] = struct{}{}
	}

	node2 := headB
	for node2 != nil {
		_, found := m[node2]
		if found {
			return node2
		}
		node2 = node2.Next
	}
	return nil
}
