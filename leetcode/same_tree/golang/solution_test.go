package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/same_tree/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		treeA    *golang.TreeNode
		treeB    *golang.TreeNode
		expected bool
	}{
		{
			name: "Example 1",
			treeA: &golang.TreeNode{
				Val:   1,
				Left:  &golang.TreeNode{Val: 2},
				Right: &golang.TreeNode{Val: 3},
			},
			treeB: &golang.TreeNode{
				Val:   1,
				Left:  &golang.TreeNode{Val: 2},
				Right: &golang.TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			name: "Example 2",
			treeA: &golang.TreeNode{
				Val:  1,
				Left: &golang.TreeNode{Val: 2},
			},
			treeB: &golang.TreeNode{
				Val:   1,
				Right: &golang.TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name: "Example 3",
			treeA: &golang.TreeNode{
				Val:   1,
				Left:  &golang.TreeNode{Val: 2},
				Right: &golang.TreeNode{Val: 1},
			},
			treeB: &golang.TreeNode{
				Val:   1,
				Left:  &golang.TreeNode{Val: 1},
				Right: &golang.TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name:     "Example 4",
			treeA:    nil,
			treeB:    nil,
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := golang.IsSameTree(tc.treeA, tc.treeB)
			require.Equal(t, tc.expected, result)
		})
	}
}
