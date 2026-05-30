package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/invert_binary_tree/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		tree     *golang.TreeNode
		expected *golang.TreeNode
	}{
		{
			name: "Example 1",
			tree: &golang.TreeNode{
				Val: 4,
				Left: &golang.TreeNode{
					Val:   2,
					Left:  &golang.TreeNode{Val: 1},
					Right: &golang.TreeNode{Val: 3},
				},
				Right: &golang.TreeNode{
					Val:   7,
					Left:  &golang.TreeNode{Val: 6},
					Right: &golang.TreeNode{Val: 9},
				},
			},
			expected: &golang.TreeNode{
				Val: 4,
				Left: &golang.TreeNode{
					Val:   7,
					Left:  &golang.TreeNode{Val: 9},
					Right: &golang.TreeNode{Val: 6},
				},
				Right: &golang.TreeNode{
					Val:   2,
					Left:  &golang.TreeNode{Val: 3},
					Right: &golang.TreeNode{Val: 1},
				},
			},
		},
		{
			name: "Example 2",
			tree: &golang.TreeNode{
				Val:   2,
				Left:  &golang.TreeNode{Val: 1},
				Right: &golang.TreeNode{Val: 3},
			},
			expected: &golang.TreeNode{
				Val:   2,
				Left:  &golang.TreeNode{Val: 3},
				Right: &golang.TreeNode{Val: 1},
			},
		},
		{
			name:     "Example 3",
			tree:     nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := golang.InvertTree(tc.tree)
			require.True(t, checkSameTree(tc.expected, result))
		})
	}
}

func checkSameTree(expected, actual *golang.TreeNode) bool {
	if expected == nil && actual == nil {
		return true
	}

	if expected == nil || actual == nil {
		return false
	}

	if expected.Val != actual.Val {
		return false
	}

	return checkSameTree(expected.Left, actual.Left) && checkSameTree(expected.Right, actual.Right)
}
