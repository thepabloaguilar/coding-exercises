package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/jump_game/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "Example 3",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Example 4",
			nums:     []int{2, 1},
			expected: true,
		},
		{
			name:     "Example 5",
			nums:     []int{2, 4, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Example 6",
			nums:     []int{1, 1, 1, 0},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.CanJump(tc.nums),
			)
		})
	}
}
