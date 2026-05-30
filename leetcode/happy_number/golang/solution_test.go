package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/happy_number/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Example 1",
			n:        19,
			expected: true,
		},
		{
			name:     "Example 2",
			n:        2,
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.IsHappy(tc.n),
			)
		})
	}
}
