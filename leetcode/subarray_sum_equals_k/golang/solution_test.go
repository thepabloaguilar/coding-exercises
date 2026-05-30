package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/subarray_sum_equals_k/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, golang.SubarraySum(tc.nums, tc.k))
		})
	}
}
