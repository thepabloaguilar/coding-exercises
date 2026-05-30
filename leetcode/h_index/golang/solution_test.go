package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/h_index/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		citations []int
		expected  int
	}{
		{
			name:      "Example 1",
			citations: []int{3, 0, 6, 1, 5},
			expected:  3,
		},
		{
			name:      "Example 2",
			citations: []int{1, 3, 1},
			expected:  1,
		},
		{
			name:      "Example 3",
			citations: []int{100},
			expected:  1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.HIndex(tc.citations),
			)
		})
	}
}
