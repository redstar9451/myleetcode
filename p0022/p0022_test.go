package p0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name   string
	n      int
	wanted []string
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"testCase2", 1, []string{"()"}},
	}

	return testCases
}

func compareArr(t *testing.T, x, y []string) {
	assert.Equal(t, len(x), len(y))
	for i := 0; i < len(x); i++ {
		assert.Equal(t, x[i], y[i])
	}
}

func TestP0022(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res := generateParenthesis(tt.n)
			compareArr(t, tt.wanted, res)
		})
	}
}
