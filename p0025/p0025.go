package p0025

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
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */
// @lc code=start

func reverseKGroup(head *ListNode, k int) *ListNode {
	reversedHead := &ListNode{Next: head}
	curGroupHead := reversedHead
	for curGroupHead != nil {
		tail := curGroupHead
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return reversedHead.Next
			}
		}
		rHead, rTail := myReverse(curGroupHead.Next, tail)
		curGroupHead.Next = rHead
		curGroupHead = rTail
	}
	return reversedHead.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	// 每次把当前节点放在end前面，指向end
	end := tail.Next
	cur := head
	for end != tail {
		tmp := cur.Next
		cur.Next = end
		end = cur
		cur = tmp
	}
	return tail, head
}

// @lc code=end
