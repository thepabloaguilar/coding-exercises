package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/symmetric_tree/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		tree     *golang.TreeNode
		expected bool
	}{
		{
			name: "Example 1",
			tree: &golang.TreeNode{
				Val: 1,
				Left: &golang.TreeNode{
					Val:   2,
					Left:  &golang.TreeNode{Val: 3},
					Right: &golang.TreeNode{Val: 4},
				},
				Right: &golang.TreeNode{
					Val:   2,
					Left:  &golang.TreeNode{Val: 4},
					Right: &golang.TreeNode{Val: 3},
				},
			},
			expected: true,
		},
		{
			name: "Example 2",
			tree: &golang.TreeNode{
				Val: 1,
				Left: &golang.TreeNode{
					Val:   2,
					Right: &golang.TreeNode{Val: 3},
				},
				Right: &golang.TreeNode{
					Val:   2,
					Right: &golang.TreeNode{Val: 3},
				},
			},
			expected: false,
		},
		{
			name:     "Example 3",
			tree:     nil,
			expected: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := golang.IsSymmetric(tc.tree)
			require.Equal(t, tc.expected, result)
		})
	}
}
