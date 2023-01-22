package main

func main() {

}

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// head = [1,2,3,4,5]
// 1. Создаем первый элемент листа
// 2.Идем по листу и каждый раз создаем новую ноду и к ней присваиваем Next весь хвост
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var currentNode = head

	newList := &ListNode{Val: currentNode.Val, Next: nil}

	for currentNode.Next != nil {
		currentNode = currentNode.Next
		newList = &ListNode{Val: currentNode.Val, Next: newList}
	}
	return newList
}
