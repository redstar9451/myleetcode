package p0328

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	for h := head; h != nil; h = h.Next {
		fmt.Printf("%v ", h.Val)
	}
	fmt.Println()
}

/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	oddHead := &ListNode{Next: nil}
	oddTail := oddHead
	evenHead := &ListNode{Next: nil}
	evenTail := evenHead
	i := 0
	for h := head; h != nil; {
		// odd
		if (i+1)%2 == 1 {
			oddTail.Next = h
			oddTail = h
		} else {
			evenTail.Next = h
			evenTail = h
		}
		tmp := h.Next
		h.Next = nil
		h = tmp
		i++
	}

	oddTail.Next = evenHead.Next
	return oddHead.Next
}

// @lc code=end
