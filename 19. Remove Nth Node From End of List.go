package main

/*
*
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
передвигаем быстрый указатель на n позиций вперед и когда fast окажется в конце
это будет значит что slow на элементе следующий за которым нужно удалить
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
