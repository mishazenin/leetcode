package main

func main() {

}

// Definition for singly-linked list.
type ListNode141 struct {
	Val  int
	Next *ListNode
}

// используем sliding window algorithm
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	var hasCycle bool
	slow := head
	fast := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			hasCycle = true
			break
		}
	}

	return hasCycle

}

// просто идем по linked list и засосываем в ключи мапы поинтеры плюс проверки
func hasCycleSlow(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	m := make(map[*ListNode]struct{})
	var currentNode = head
	for {
		if currentNode.Next == nil || currentNode == nil {
			return false
		}
		currentNode = currentNode.Next
		if _, found := m[currentNode.Next]; found {
			return true
		}
		m[currentNode] = struct{}{}
	}
}
