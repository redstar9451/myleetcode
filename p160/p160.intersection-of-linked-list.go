package p160

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hA := headA
	hB := headB
	var match *ListNode
	for {
		if hA == hB {
			match = hA
			break
		}
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}

		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}
	return match
}

// @lc code=end
