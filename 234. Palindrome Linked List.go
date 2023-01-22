package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
	Все просто

1. делаем стэк значений
2.Идем с разных сторон и проверяем равенство
*/
func isPalindromeList(head *ListNode) bool {

	if head == nil {
		return true
	}

	var data []int
	for cur := head; cur != nil; cur = cur.Next {
		data = append(data, cur.Val)
	}

	for i, j := 0, len(data)-1; i <= j; {
		if data[i] != data[j] {
			return false
		}
		i++
		j--
	}
	return true
}
