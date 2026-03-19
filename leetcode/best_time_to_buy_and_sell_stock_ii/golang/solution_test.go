package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/best_time_to_buy_and_sell_stock_ii/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "Example 2",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Example 3",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.MaxProfit(tc.prices),
			)
		})
	}
}
