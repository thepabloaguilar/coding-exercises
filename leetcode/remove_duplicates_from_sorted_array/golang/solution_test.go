package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/remove_duplicates_from_sorted_array/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			remainingNums := golang.RemoveDuplicates(tc.nums)

			require.Equal(t, len(tc.expected), remainingNums)
			require.ElementsMatch(t, tc.expected, tc.nums[:remainingNums])
		})
	}
}
