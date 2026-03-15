package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/binary_tree_level_order_traversal_ii/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		tree     *golang.TreeNode
		expected [][]int
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
			expected: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			name:     "Example 2",
			tree:     &golang.TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			name:     "Example 3",
			tree:     nil,
			expected: [][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := golang.LevelOrderBottom(tc.tree)
			require.Equal(t, tc.expected, result)
		})
	}
}
