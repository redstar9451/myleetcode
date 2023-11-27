package p21

import (
	"sort"
	"testing"
)

type TT struct {
	name  string
	listA []int
	listB []int

	a *ListNode // a 链表
	b *ListNode // b 链表
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 3, 5}, []int{2, 4, 6}, nil, nil},
		{"testCase2", []int{1, 2, 3, 4, 5}, []int{6, 7, 8}, nil, nil},
		{"testCase3", []int{}, []int{}, nil, nil},
	}

	for i, testCase := range testCases {
		testCase.a = convert2List(testCase.listA)
		testCase.b = convert2List(testCase.listB)
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

func TestP21(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.a, tt.b)
			gotList := []int{}
			for node := got; node != nil; node = node.Next {
				gotList = append(gotList, node.Val)
			}
			if len(gotList) != len(tt.listA)+len(tt.listB) {
				t.Errorf("Length of merged list is not equal :w to the sum of lengths of the two lists")
			}
			mergedList := append(tt.listA, tt.listB...)
			sort.Ints(mergedList)
			for i := range gotList {
				if gotList[i] != mergedList[i] {
					t.Errorf("gotList and mergedList are not equal")
				}
			}
		})
	}
}
