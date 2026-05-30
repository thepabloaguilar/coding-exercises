package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/number_of_provinces/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          [][]int
		expectedOutput int
	}{
		{
			name:           "Example 1",
			input:          [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			expectedOutput: 2,
		},
		{
			name:           "Example 2",
			input:          [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			expectedOutput: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expectedOutput,
				golang.FindCircleNum(tc.input),
			)
		})
	}
}
