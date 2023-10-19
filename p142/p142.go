// package p142

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}
		slow = slow.Next
		if fast == slow {
			for h := head; h != slow; {
				h = h.Next
				slow = slow.Next
			}
			return slow
		}
	}
}

// @lc code=end
