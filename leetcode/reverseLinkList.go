package leetcode
//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL
type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := head.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return head
}
