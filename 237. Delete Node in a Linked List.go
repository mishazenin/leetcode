package main

// хз просто работает
// так как нужно удалить значение а не сам поинтер(ноду)
// то мы просто вставляем значение из след ноды а next ставим next next
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
