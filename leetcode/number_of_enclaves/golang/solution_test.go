package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/number_of_enclaves/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name           string
		input          [][]int
		expectedOutput int
	}{
		{
			name:           "Example 1",
			input:          [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			expectedOutput: 3,
		},
		{
			name:           "Example 2",
			input:          [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
			expectedOutput: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expectedOutput,
				golang.NumEnclaves(tc.input),
			)
		})
	}
}
