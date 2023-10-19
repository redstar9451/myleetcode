package p86

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	h := &ListNode{0, head}
	lessHead := &ListNode{0, nil} // head of the chain which elements are less than x
	lessTail := lessHead
	var newHead *ListNode

	pre, cur := h, h.Next
	for cur != nil {
		if cur.Val < x {
			pre.Next = cur.Next
			lessTail.Next = cur
			lessTail = cur
		} else {
			if newHead == nil {
				newHead = cur
			}
			pre = pre.Next
		}
		cur = cur.Next
	}
	lessTail.Next = newHead
	return lessHead.Next
}

// @lc code=end
