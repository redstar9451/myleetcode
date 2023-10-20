package p138

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name string
	list []int

	head *Node
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 4, 3, 2, 5, 2}, nil},
		{"testCase2", []int{1, 4, 3, 2, 5, 2}, nil},
		{"testCase3", []int{1}, nil},
	}
	for i, v := range testCases {
		v.head = convert2List(v.list)
		testCases[i] = v
	}
	return testCases
}

func convert2List(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	listNode := make([]*Node, len(values))
	head := &Node{Val: values[0]}
	listNode[0] = head
	current := head
	for i, v := range values[1:] {
		current.Next = &Node{Val: v}
		current = current.Next
		listNode[i] = current
	}

	for cur := head; cur != nil; cur = cur.Next {
		random := rand.Intn(len(values))
		if random == 0 {
			cur.Random = nil
		} else {
			cur.Random = listNode[random]
		}
	}
	return head
}

func dumpArray(head *Node) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}

func dumpRandomArray(head *Node) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Random != nil {
			result = append(result, cur.Random.Val)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func TestP138(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			dumpArray(tt.head)
			dumpRandomArray(tt.head)
			want := copyRandomList(tt.head)
			wantCur, inputCur := want, tt.head
			for wantCur != nil && inputCur != nil {
				assert.Equal(t, wantCur.Val, inputCur.Val)
				if wantCur.Random != nil {
					assert.NotEmpty(t, inputCur.Random)
					assert.Equal(t, wantCur.Random.Val, inputCur.Random.Val)
				}
				wantCur = wantCur.Next
				inputCur = inputCur.Next
			}
			assert.Nil(t, wantCur)
			assert.Nil(t, inputCur)
		})
	}
}
