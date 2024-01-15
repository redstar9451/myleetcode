package p0739

import "testing"

type TT struct {
	name   string
	input  []int
	output []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"testCase2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"testCase3", []int{30, 60, 90}, []int{1, 1, 0}},
	}

	return testCases
}

func TestP946(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res := dailyTemperatures(tt.input)
			if len(tt.output) != len(res) {
				t.Errorf("output %v, want %v", res, tt.output)
				t.FailNow()
			}
			for i, v := range res {
				if v != tt.output[i] {
					t.Errorf("output %v, want %v", res, tt.output)
					t.FailNow()
				}
			}

		})
	}
}
