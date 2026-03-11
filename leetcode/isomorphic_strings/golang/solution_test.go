package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/isomorphic_strings/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "egg",
			t:        "add",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "f11",
			t:        "b23",
			expected: false,
		},
		{
			name:     "Example 3",
			s:        "paper",
			t:        "title",
			expected: true,
		},
		{
			name:     "Example 4",
			s:        "badc",
			t:        "baba",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.IsIsomorphic(tc.s, tc.t),
			)
		})
	}
}
