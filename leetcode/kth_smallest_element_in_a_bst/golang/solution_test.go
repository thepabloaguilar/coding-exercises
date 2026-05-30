package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/kth_smallest_element_in_a_bst/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		root     *golang.TreeNode
		k        int
		expected int
	}{
		{
			name: "Example 1",
			root: &golang.TreeNode{
				Val: 3,
				Left: &golang.TreeNode{
					Val:   1,
					Right: &golang.TreeNode{Val: 2},
				},
				Right: &golang.TreeNode{Val: 4},
			},
			k:        1,
			expected: 1,
		},
		{
			name: "Example 2",
			root: &golang.TreeNode{
				Val: 5,
				Left: &golang.TreeNode{
					Val: 3,
					Left: &golang.TreeNode{
						Val:  2,
						Left: &golang.TreeNode{Val: 1},
					},
					Right: &golang.TreeNode{Val: 4},
				},
				Right: &golang.TreeNode{Val: 6},
			},
			k:        3,
			expected: 3,
		},
		{
			name: "Example 3",
			root: &golang.TreeNode{
				Val: 5,
				Left: &golang.TreeNode{
					Val: 3,
					Left: &golang.TreeNode{
						Val:  2,
						Left: &golang.TreeNode{Val: 1},
					},
					Right: &golang.TreeNode{Val: 4},
				},
				Right: &golang.TreeNode{Val: 6},
			},
			k:        4,
			expected: 4,
		},
		{
			name:     "Example 4",
			root:     &golang.TreeNode{Val: 5},
			k:        1,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := golang.KthSmallest(tc.root, tc.k)
			require.Equal(t, tc.expected, result)
		})
	}
}
