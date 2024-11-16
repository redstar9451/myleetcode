package p0170

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name   string
	input  []int
	wanted int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{9, 7, 5, 4, 6}, 8},
		{"testCase2", []int{7, 5, 6, 4}, 5},
	}
	return testCases
}

func TestP0170(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res := reversePairs(tt.input)
			assert.Equal(t, tt.wanted, res)
		})
	}
}
