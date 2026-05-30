package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/gas_station/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "Example 3",
			gas:      []int{3, 3, 4},
			cost:     []int{3, 4, 4},
			expected: -1,
		},
		{
			name:     "Example 4",
			gas:      []int{5, 1, 2, 3, 4},
			cost:     []int{4, 4, 1, 5, 1},
			expected: 4,
		},
		{
			name:     "Example 5",
			gas:      []int{5, 8, 2, 8},
			cost:     []int{6, 5, 6, 6},
			expected: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(
				t,
				tc.expected,
				golang.CanCompleteCircuit(tc.gas, tc.cost),
			)
		})
	}
}
