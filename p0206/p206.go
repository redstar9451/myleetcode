package p206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 现把头节点放在尾部，此时头节点的next指向nil
	tail := head.Next
	head.Next = nil

	newHead := head
	// 每次把当前节点放在头部，然后把当前节点的next指向新的头部
	for tail != nil {
		tmp := tail.Next
		tail.Next = newHead
		newHead = tail
		tail = tmp
	}
	return newHead
}
