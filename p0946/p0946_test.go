package p0946

import "testing"

type TT struct {
	name   string
	pushed []int
	poped  []int
	want   bool
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{"testCase2", []int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	}

	return testCases
}

func TestP946(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res := validateStackSequences(tt.pushed, tt.poped)
			if res != tt.want {
				t.Errorf("%v %v", res, tt.want)
				t.FailNow()
			}
		})
	}
}
