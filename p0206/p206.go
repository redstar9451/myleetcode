package p206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tail := head.Next
	head.Next = nil
	newHead := head
	for tail != nil {
		tmp := tail.Next
		tail.Next = newHead
		newHead = tail
		tail = tmp
	}
	return newHead
}
