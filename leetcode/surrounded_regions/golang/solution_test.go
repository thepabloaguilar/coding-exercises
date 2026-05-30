package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/surrounded_regions/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		board    [][]byte
		expected [][]byte
	}{
		{
			name: "Example 1",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name:     "Example 2",
			board:    [][]byte{{'X'}},
			expected: [][]byte{{'X'}},
		},
		{
			name: "Example 3",
			board: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			expected: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "Example 4",
			board: [][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
			expected: [][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
		},
		{
			name: "Example 5",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			},
			expected: [][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
			},
		},
		{
			name: "Example 6",
			board: [][]byte{
				{'X', 'X', 'X', 'X', 'O'},
				{'X', 'O', 'O', 'X', 'X'},
				{'O', 'X', 'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X', 'O'},
				{'X', 'X', 'X', 'X', 'X'},
				{'O', 'X', 'X', 'X', 'X'},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			golang.Solve(tc.board)
			require.Equal(t, tc.expected, tc.board)
		})
	}
}
