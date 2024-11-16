package p0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name string
	nums []int
	val  int
	want int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{3, 2, 2, 3}, 3, 2},
		{"testCase2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	return testCases
}

func TestP0027(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res := removeElement(tt.nums, tt.val)
			assert.Equal(t, tt.want, res, "test failed")
		})
	}
}
