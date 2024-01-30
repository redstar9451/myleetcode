package p0234

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name   string
	listA  []int
	wanted bool
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5}, false},
		{"testCase2", []int{1, 2, 2, 1}, true},
		{"testCase2", []int{1, 2, 3, 2, 1}, true},
	}

	return testCases
}

func convert2List(values []int) *ListNode {
	var head *ListNode
	var current *ListNode
	for _, val := range values {
		if head == nil {
			head = &ListNode{Val: val}
			current = head
		} else {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
	}
	return head
}

func convert2Arr(head *ListNode) []int {
	res := make([]int, 0)
	for h := head; h != nil; h = h.Next {
		res = append(res, h.Val)
	}
	return res
}

func TestP25(t *testing.T) {
	testcases := generateTestCases()
	printList(nil)
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			input := convert2List(tt.listA)
			res := isPalindrome(input)
			assert.Equal(t, res, tt.wanted)
		})
	}
}
