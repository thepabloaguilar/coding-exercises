package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/find_the_index_of_the_first_occurrence_in_a_string/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Example 1",
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			name:     "Example 2",
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			name:     "Example 3",
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
		{
			name:     "Example 4",
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
		{
			name:     "Example 5",
			haystack: "aaa",
			needle:   "aaaa",
			expected: -1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.StrStr(tc.haystack, tc.needle),
			)
		})
	}
}
