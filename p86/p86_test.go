package p86

import "testing"

type TT struct {
	name string
	list []int
	x    int

	head *ListNode

	want []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 4, 3, 2, 5, 2}, 3, nil, []int{1, 2, 2, 4, 3, 5}},
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

func TestP86(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.head, tt.x)
			i := 0
			for n := got; n != nil; n = n.Next {
				if n.Val != tt.want[i] {
					t.Errorf("partition() = %v, want %v", n.Val, tt.want[i])
				}
				i++
			}
		})
	}
}
