package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// [0,3,1,0,4,5,2,0]

func mergeNodes(head *ListNode) *ListNode {
	currNode := head.Next

	var (
		newList  *ListNode
		listHead *ListNode
		sum      int
	)

	for currNode != nil {

		if currNode.Val == 0 && sum != 0 {

			if newList == nil {
				newList = &ListNode{Val: sum}
				listHead = newList
				sum = 0
				continue
			}

			newList.Next = &ListNode{Val: sum}
			newList = newList.Next
			sum = 0
		}

		sum += currNode.Val
		currNode = currNode.Next
	}

	return listHead
}
