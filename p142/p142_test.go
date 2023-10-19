package p142

import "testing"

type TT struct {
	name     string
	list     []int
	position int

	head      *ListNode
	loopPoint *ListNode
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 4, 3, 2, 5, 2}, 3, nil, nil},
		{"testCase2", []int{1, 4, 3, 2, 5, 2}, -1, nil, nil},
		{"testCase3", []int{1}, -1, nil, nil},
	}
	for i, v := range testCases {
		v.head, v.loopPoint = convert2List(v.list, v.position)
		testCases[i] = v
	}
	return testCases
}

func convert2List(values []int, position int) (*ListNode, *ListNode) {
	var head *ListNode
	var current *ListNode
	var loopPoint *ListNode
	for i, val := range values {
		if head == nil {
			head = &ListNode{Val: val}
			current = head
		} else {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		if i == position {
			loopPoint = current
		}
	}
	if loopPoint != nil {
		current.Next = loopPoint
	}
	return head, loopPoint
}

func TestP142(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			want := detectCycle(tt.head)
			if want != tt.loopPoint {
				t.Errorf("got %v, want %v", want, tt.loopPoint)
			}

		})
	}
}
