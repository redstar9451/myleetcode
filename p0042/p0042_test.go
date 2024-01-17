package p042

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name  string
	input []int
	want  int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"testCase2", []int{4, 2, 0, 3, 2, 5}, 9},
	}

	return testCases
}

func TestP042(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res := trap(tt.input)
			assert.Equal(t, tt.want, res, "test failed")
		})
	}
}
