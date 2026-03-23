package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/longest_consecutive_sequence/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 0, 1, 2},
			expected: 3,
		},
		{
			name:     "Example 4",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "Example 5",
			nums:     []int{0, -1},
			expected: 2,
		},
		{
			name:     "Example 6",
			nums:     make([]int, 0),
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.LongestConsecutive(tc.nums),
			)
		})
	}
}
