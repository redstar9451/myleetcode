package p0203

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start

func removeElements(head *ListNode, val int) *ListNode {
	temp := &ListNode{0, head}
	h := temp
	for h.Next != nil {
		if h.Next.Val == val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	return temp.Next
}

// @lc code=end
