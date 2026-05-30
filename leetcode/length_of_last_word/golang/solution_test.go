package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/length_of_last_word/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "Hello World",
			expected: 5,
		},
		{
			name:     "Example 2",
			s:        "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Example 3",
			s:        "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "Example 4",
			s:        "Hello",
			expected: 5,
		},
		{
			name:     "Example 5",
			s:        "a ",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.LengthOfLastWord(tc.s),
			)
		})
	}
}
