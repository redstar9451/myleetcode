package p86

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{0, head}
	pLeft := dummyHead
	for i := 0; i < left-1; i++ {
		pLeft = pLeft.Next
	}
	// the first node of the chain which will be reversed
	cur := pLeft.Next
	for i := left - 1; i < right-1; i++ {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		tmp.Next = pLeft.Next
		pLeft.Next = tmp
	}
	return dummyHead.Next
}

// @lc code=end
