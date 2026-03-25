package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/house_robber_iii/golang"
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
				Val: 3,
				Left: &golang.TreeNode{
					Val:   2,
					Right: &golang.TreeNode{Val: 3},
				},
				Right: &golang.TreeNode{
					Val:   3,
					Right: &golang.TreeNode{Val: 1},
				},
			},
			expected: 7,
		},
		{
			name: "Example 2",
			tree: &golang.TreeNode{
				Val: 3,
				Left: &golang.TreeNode{
					Val:   4,
					Left:  &golang.TreeNode{Val: 1},
					Right: &golang.TreeNode{Val: 3},
				},
				Right: &golang.TreeNode{
					Val:   5,
					Right: &golang.TreeNode{Val: 1},
				},
			},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, golang.Rob(tc.tree))
		})
	}
}
