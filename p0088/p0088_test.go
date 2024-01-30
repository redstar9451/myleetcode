package p0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TT struct {
	name  string
	nums1 []int
	m     int
	nums2 []int
	n     int
	want  []int
}

func generateTestCases() []TT {
	testCases := []TT{
		{"testCase1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"testCase2", []int{1}, 1, []int{}, 0, []int{1}},
	}

	return testCases
}

func compareArr(t *testing.T, x, y []int) {
	assert.Equal(t, len(x), len(y))
	for i := 0; i < len(x); i++ {
		assert.Equal(t, x[i], y[i])
	}
}

func TestP0088(t *testing.T) {
	testcases := generateTestCases()
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			compareArr(t, tt.want, tt.nums1)
		})
	}
}
