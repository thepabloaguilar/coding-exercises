package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/is_subsequence/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
		{
			name:     "Example 3",
			s:        "",
			t:        "ahbgdc",
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
				golang.IsSubsequence(tc.s, tc.t),
			)
		})
	}
}
