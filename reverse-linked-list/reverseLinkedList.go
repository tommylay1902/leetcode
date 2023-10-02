

func reverseList(l *ListNode) *ListNode {
	var prev *ListNode
	prev = nil
	curr = l
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}

	return prev
}
