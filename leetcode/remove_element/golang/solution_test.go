package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/remove_element/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		val      int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: []int{2, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: []int{0, 1, 4, 0, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			remainingNums := golang.RemoveElement(tc.nums, tc.val)

			require.Equal(t, len(tc.expected), remainingNums)
			require.ElementsMatch(t, tc.expected, tc.nums[:remainingNums])
		})
	}
}
