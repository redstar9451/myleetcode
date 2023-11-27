package p20

import (
	"testing"
)

type TT struct {
	name   string
	input  string
	result bool
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", "(){}[]", true},
		{"testCase1", "({}[]", false},
	}
	return testCases
}

func TestP20(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			if got != tt.result {
				t.Errorf("isValid(%v) = %v; want %v", tt.input, got, tt.result)
			}
		})
	}
}
