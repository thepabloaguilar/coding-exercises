package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/count_complete_tree_nodes/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		root     *golang.TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &golang.TreeNode{
				Val: 1,
				Left: &golang.TreeNode{
					Val:   2,
					Left:  &golang.TreeNode{Val: 4},
					Right: &golang.TreeNode{Val: 5},
				},
				Right: &golang.TreeNode{
					Val:  3,
					Left: &golang.TreeNode{Val: 6},
				},
			},
			expected: 6,
		},
		{
			name:     "Example 2",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Example 3",
			root:     &golang.TreeNode{Val: 1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := golang.CountNodes(tc.root)
			require.Equal(t, tc.expected, result)
		})
	}
}
