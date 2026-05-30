package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/valid_parentheses/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "()",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "(]",
			expected: false,
		},
		{
			name:     "Example 4",
			s:        "([])",
			expected: true,
		},
		{
			name:     "Example 5",
			s:        "([)]",
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
				golang.IsValid(tc.s),
			)
		})
	}
}
