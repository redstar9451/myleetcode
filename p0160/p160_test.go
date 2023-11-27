package p160

import "testing"

type TT struct {
	name  string
	listA []int
	listB []int
	listC []int // 是否相交

	a    *ListNode // a 链表
	b    *ListNode // b 链表
	want *ListNode // 相交的节点
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5}, []int{6, 7, 8}, []int{9, 10, 11}, nil, nil, nil},
		{"testCase2", []int{1, 2, 3, 4, 5}, []int{6, 7, 8}, []int{}, nil, nil, nil},
		{"testCase3", []int{1}, []int{2}, []int{}, nil, nil, nil},
		{"testCase4", []int{}, []int{}, []int{}, nil, nil, nil},
	}

	for i, testCase := range testCases {
		testCase.a = convert2List(testCase.listA)
		testCase.b = convert2List(testCase.listB)
		c := convert2List(testCase.listC)
		testCase.a = appendList(testCase.a, c)
		testCase.b = appendList(testCase.b, c)
		testCase.want = c
		testCases[i] = testCase
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

func appendList(x, y *ListNode) *ListNode {
	if x == nil {
		x = y
	} else {
		current := x
		for current.Next != nil {
			current = current.Next
		}
		current.Next = y
	}
	return x
}

func TestP106(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := getIntersectionNode(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
