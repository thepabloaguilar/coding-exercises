package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/roman_to_integer/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "III",
			expected: 3,
		},
		{
			name:     "Example 2",
			s:        "LVIII",
			expected: 58,
		},
		{
			name:     "Example 3",
			s:        "MCMXCIV",
			expected: 1994,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.RomanToInt(tc.s),
			)
		})
	}
}
