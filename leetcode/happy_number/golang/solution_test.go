package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/happy_number/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Example 1",
			n:        19,
			expected: true,
		},
		{
			name:     "Example 2",
			n:        2,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.IsHappy(tc.n),
			)
		})
	}
}
