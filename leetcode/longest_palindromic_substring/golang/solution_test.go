package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/longest_palindromic_substring/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Example 1",
			s:        "babad",
			expected: "bab",
		},
		{
			name:     "Example 2",
			s:        "cbbd",
			expected: "bb",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.LongestPalindrome(tc.s),
			)
		})
	}
}
