package p75

import (
	"reflect"
	"sort"
	"testing"
)

type TT struct {
	name  string
	input []int
	want  []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{0, 2, 1}, nil},
		{"testCase2", []int{0, 2, 1, 0, 1}, nil},
		{"testCase3", []int{2, 1, 0, 0, 2, 1, 1}, nil},
		{"testCase4", []int{1, 2, 0}, nil},
	}
	for i, testCase := range testCases {
		testCase.want = make([]int, len(testCase.input))
		copy(testCase.want, testCase.input)
		sort.Ints(testCase.want)
		testCases[i] = testCase
	}

	return testCases
}

func TestP75(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
