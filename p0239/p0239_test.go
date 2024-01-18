package p0239

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name string
	arr  []int
	k    int
	want []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"testCase2", []int{1}, 1, []int{1}},
		{"testCase3", []int{4, 3, 11}, 3, []int{11}},
	}

	return testCases
}

func TestP0239(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res := maxSlidingWindow(tt.arr, tt.k)
			assert.Equal(t, len(tt.want), len(res))
			for i, v := range tt.want {
				assert.Equal(t, v, res[i])
			}
		})
	}
}
