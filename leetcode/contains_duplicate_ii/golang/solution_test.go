package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/contains_duplicate_ii/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 0, 1, 1},
			k:        1,
			expected: true,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3, 1, 2, 3},
			k:        2,
			expected: false,
		},
		{
			name:     "Example 4",
			nums:     []int{1, 2, 3, 1, 2, 3, 2},
			k:        2,
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.ContainsNearbyDuplicate(tc.nums, tc.k),
			)
		})
	}
}
