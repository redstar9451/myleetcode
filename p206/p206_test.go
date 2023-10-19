package p206

import "testing"

type TT struct {
	name      string
	input     []int
	inputList *ListNode
	want      []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5}, nil, nil},
		{"testCase2", []int{1}, nil, nil},
		{"testCase2", []int{}, nil, nil},
	}

	for i, testCase := range testCases {
		testCase.inputList = convert2List(testCase.input)
		testCase.want = reverseInput(testCase.input)
		testCases[i] = testCase
	}
	return testCases
}

func reverseInput(input []int) []int {
	reversed := make([]int, len(input))
	for i, val := range input {
		reversed[len(input)-1-i] = val
	}
	return reversed
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

func TestP106(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			output := reverseList(tt.inputList)

			outputArr := make([]int, 0, len(tt.input))
			for output != nil {
				outputArr = append(outputArr, output.Val)
				output = output.Next
			}
			if len(tt.input) != len(outputArr) {
				t.Errorf("input %v does not equal with output %v", tt.input, outputArr)
				t.FailNow()
			}
			for i, val := range tt.want {
				if val != outputArr[i] {
					t.Errorf("input %v does not equal with output %v", tt.input, outputArr)
					t.FailNow()
					break
				}
			}
		})
	}
}
