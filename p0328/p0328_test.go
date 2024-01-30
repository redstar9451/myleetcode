package p0328

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name   string
	listA  []int
	wanted []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{"testCase2", []int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
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

func compareArr(t *testing.T, x, y []int) {
	assert.Equal(t, len(x), len(y))
	for i := 0; i < len(x); i++ {
		assert.Equal(t, x[i], y[i])
	}
}

func TestP25(t *testing.T) {
	testcases := generateTestCases()
	printList(nil)
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			input := convert2List(tt.listA)
			res := oddEvenList(input)
			resArr := convert2Arr(res)
			compareArr(t, tt.wanted, resArr)
		})
	}
}
