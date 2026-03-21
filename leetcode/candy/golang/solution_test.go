package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/candy/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "Example 1",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "Example 2",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			name:     "Example 3",
			ratings:  []int{1, 2, 87, 87, 87, 2, 1},
			expected: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.Candy(tc.ratings),
			)
		})
	}
}
