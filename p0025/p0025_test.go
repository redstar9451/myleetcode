package p0025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name  string
	listA []int
	k     int
	listB []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"testCase2", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
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
			res := reverseKGroup(input, tt.k)
			output := convert2Arr(res)
			assert.Equal(t, len(output), len(tt.listB))
			for i := 0; i < len(output); i++ {
				assert.Equal(t, output[i], tt.listB[i])
			}
		})
	}
}
