package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/word_break/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "Example 4",
			s:        "cars",
			wordDict: []string{"car", "ca", "rs"},
			expected: true,
		},
		{
			name:     "Example 5",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.WordBreak(tc.s, tc.wordDict),
			)
		})
	}
}
