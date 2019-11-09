package leetcode

// 206. Reverse Linked List
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	q := head
	p := head.Next
	for p != nil {
		q1 := p
		p1 := p.Next
		p.Next = q
		p = p1
		q = q1
	}
	head.Next = nil
	return q
}
