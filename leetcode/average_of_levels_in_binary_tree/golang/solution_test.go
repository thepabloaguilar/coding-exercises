package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/average_of_levels_in_binary_tree/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		tree     *golang.TreeNode
		expected []float64
	}{
		{
			name: "Example 1",
			tree: &golang.TreeNode{
				Val:  3,
				Left: &golang.TreeNode{Val: 9},
				Right: &golang.TreeNode{
					Val:   20,
					Left:  &golang.TreeNode{Val: 15},
					Right: &golang.TreeNode{Val: 7},
				},
			},
			expected: []float64{3, 14.5, 11},
		},
		{
			name: "Example 2",
			tree: &golang.TreeNode{
				Val: 3,
				Left: &golang.TreeNode{
					Val:   9,
					Left:  &golang.TreeNode{Val: 15},
					Right: &golang.TreeNode{Val: 7},
				},
				Right: &golang.TreeNode{Val: 20},
			},
			expected: []float64{3, 14.5, 11},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := golang.AverageOfLevels(tc.tree)
			require.Equal(t, tc.expected, result)
		})
	}
}
