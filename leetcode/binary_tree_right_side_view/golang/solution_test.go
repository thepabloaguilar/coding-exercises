package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/binary_tree_right_side_view/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		tree     *golang.TreeNode
		expected []int
	}{
		{
			name: "Example 1",
			tree: &golang.TreeNode{
				Val: 1,
				Left: &golang.TreeNode{
					Val:   2,
					Right: &golang.TreeNode{Val: 5},
				},
				Right: &golang.TreeNode{
					Val:   3,
					Right: &golang.TreeNode{Val: 4},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name: "Example 2",
			tree: &golang.TreeNode{
				Val: 1,
				Left: &golang.TreeNode{
					Val: 2,
					Left: &golang.TreeNode{
						Val:  4,
						Left: &golang.TreeNode{Val: 5},
					},
				},
				Right: &golang.TreeNode{Val: 3},
			},
			expected: []int{1, 3, 4, 5},
		},
		{
			name: "Example 3",
			tree: &golang.TreeNode{
				Val:   1,
				Right: &golang.TreeNode{Val: 3},
			},
			expected: []int{1, 3},
		},
		{
			name:     "Example 4",
			tree:     nil,
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := golang.RightSideView(tc.tree)
			require.Equal(t, tc.expected, result)
		})
	}
}
