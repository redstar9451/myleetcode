package p0234

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
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 转成数组，比较是否回文
	arr := make([]int, 0)
	for h := head; h != nil; h = h.Next {
		arr = append(arr, h.Val)
	}
	res := true
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			res = false
			break
		}
	}
	return res
}

// @lc code=end
