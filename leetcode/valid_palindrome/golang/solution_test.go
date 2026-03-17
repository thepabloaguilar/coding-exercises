package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/valid_palindrome/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "Example 3",
			s:        " ",
			expected: true,
		},
		{
			name:     "Example 4",
			s:        ",,,,,,,,,,,,acva",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.IsPalindrome(tc.s),
			)
		})
	}
}
