package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/path_sum/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		root      *golang.TreeNode
		targetSum int
		expected  bool
	}{
		{
			name: "Example 1",
			root: &golang.TreeNode{
				Val: 5,
				Left: &golang.TreeNode{
					Val: 4,
					Left: &golang.TreeNode{
						Val:   11,
						Left:  &golang.TreeNode{Val: 7},
						Right: &golang.TreeNode{Val: 2},
					},
				},
				Right: &golang.TreeNode{
					Val:  8,
					Left: &golang.TreeNode{Val: 13},
					Right: &golang.TreeNode{
						Val:   4,
						Right: &golang.TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			expected:  true,
		},
		{
			name: "Example 2",
			root: &golang.TreeNode{
				Val:   1,
				Left:  &golang.TreeNode{Val: 2},
				Right: &golang.TreeNode{Val: 3},
			},
			targetSum: 5,
			expected:  false,
		},
		{
			name:      "Example 3",
			root:      nil,
			targetSum: 0,
			expected:  false,
		},
		{
			name: "Example 4",
			root: &golang.TreeNode{
				Val:  1,
				Left: &golang.TreeNode{Val: 2},
			},
			targetSum: 1,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := golang.HasPathSum(tc.root, tc.targetSum)
			require.Equal(t, tc.expected, result)
		})
	}
}
