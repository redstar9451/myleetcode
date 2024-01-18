package p0203

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name  string
	head  []int
	value int
	want  []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{"testCase2", []int{1, 2, 3, 4, 5, 6}, 3, []int{1, 2, 4, 5, 6}},
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
	arr := make([]int, 0)
	for h := head; h != nil; h = h.Next {
		arr = append(arr, h.Val)
	}
	return arr
}

func TestP0203(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			list := convert2List(tt.head)
			res := removeElements(list, tt.value)
			output := convert2Arr(res)
			head := []int{}
			for _, v := range tt.head {
				if v != tt.value {
					head = append(head, v)
				}
			}
			assert.Equal(t, len(head), len(output))
			for i, v := range head {
				assert.Equal(t, v, output[i])
			}
		})
	}
}
