package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/jump_game_ii/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{2, 1},
			expected: 1,
		},
		{
			name:     "Example 4",
			nums:     []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "Example 5",
			nums:     []int{3, 2, 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.Jump(tc.nums),
			)
		})
	}
}
