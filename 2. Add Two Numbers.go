package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
     		2, 4, 3
			5, 6, 4
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList *ListNode
	var dept int
	for l1.Next != nil && l2.Next != nil {
		nodeSum := getSum(l1, l2, dept)
		if nodeSum > 9 {
			dept = 1
			nodeSum = nodeSum % 10
		}
		dept = dept
		if l1.Next != nil {
			l1 = l1.Next
		}
		if l2.Next != nil {
			l2 = l2.Next
		}
		if newList == nil {
			newList = &ListNode{
				Val:  nodeSum,
				Next: nil,
			}
			continue
		}
		newList.Next = &ListNode{
			Val:  nodeSum,
			Next: nil,
		}
	}
	nodeSum := getSum(l1, l2, dept)
	if nodeSum > 9 {
		dept = 1
		if nodeSum != 10 {
			nodeSum = nodeSum % 10
		} else {
			nodeSum = 0
		}
	}
	newList.Next = &ListNode{
		Val:  nodeSum,
		Next: nil,
	}
	return newList
}
func getSum(n1, n2 *ListNode, dept int) int {
	return n1.Val + n2.Val + dept
}
