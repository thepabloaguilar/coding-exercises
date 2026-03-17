package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/maximum_depth_of_binary_tree/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		tree     *golang.TreeNode
		expected int
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
			expected: 3,
		},
		{
			name: "Example 2",
			tree: &golang.TreeNode{
				Val:   1,
				Right: &golang.TreeNode{Val: 2},
			},
			expected: 2,
		},
		{
			name:     "Example 3",
			tree:     nil,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := golang.MaxDepth(tc.tree)
			require.Equal(t, tc.expected, result)
		})
	}
}
