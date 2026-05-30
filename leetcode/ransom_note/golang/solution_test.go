package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/ransom_note/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		ransomNote string
		magazine   string
		expected   bool
	}{
		{
			name:       "Example 1",
			ransomNote: "a",
			magazine:   "b",
			expected:   false,
		},
		{
			name:       "Example 2",
			ransomNote: "aa",
			magazine:   "ab",
			expected:   false,
		},
		{
			name:       "Example 3",
			ransomNote: "aa",
			magazine:   "aab",
			expected:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.CanConstruct(tc.ransomNote, tc.magazine),
			)
		})
	}
}
