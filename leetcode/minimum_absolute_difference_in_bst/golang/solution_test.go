package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/minimum_absolute_difference_in_bst/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		root     *golang.TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &golang.TreeNode{
				Val: 4,
				Left: &golang.TreeNode{
					Val:   2,
					Left:  &golang.TreeNode{Val: 1},
					Right: &golang.TreeNode{Val: 3},
				},
				Right: &golang.TreeNode{Val: 6},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			root: &golang.TreeNode{
				Val:  1,
				Left: &golang.TreeNode{Val: 0},
				Right: &golang.TreeNode{
					Val:   48,
					Left:  &golang.TreeNode{Val: 12},
					Right: &golang.TreeNode{Val: 49},
				},
			},
			expected: 1,
		},
		{
			name: "Example 3",
			root: &golang.TreeNode{
				Val: 1,
				Right: &golang.TreeNode{
					Val:  5,
					Left: &golang.TreeNode{Val: 3},
				},
			},
			expected: 2,
		},
		{
			name: "Example 4",
			root: &golang.TreeNode{
				Val: 236,
				Left: &golang.TreeNode{
					Val:   104,
					Right: &golang.TreeNode{Val: 227},
				},
				Right: &golang.TreeNode{
					Val:   701,
					Right: &golang.TreeNode{Val: 911},
				},
			},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(
				t,
				tc.expected,
				golang.GetMinimumDifference(tc.root),
			)
		})
	}
}
