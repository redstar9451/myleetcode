package p86

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name  string
	list  []int
	left  int
	right int

	head *ListNode

	want []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5}, 2, 4, nil, []int{1, 4, 3, 2, 5}},
	}
	for i, v := range testCases {
		v.head = convert2List(v.list)
		testCases[i] = v
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

func TestP92(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			wantList := reverseBetween(tt.head, tt.left, tt.right)
			want := make([]int, len(tt.want))
			i := 0
			for n := wantList; n != nil; n = n.Next {
				want[i] = n.Val
				i++
			}
			assert.Equal(t, len(want), len(tt.want))
			for i, v := range tt.want {
				assert.Equal(t, v, want[i])
			}
		})
	}
}
