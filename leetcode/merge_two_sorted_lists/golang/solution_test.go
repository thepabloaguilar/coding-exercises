package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/merge_two_sorted_lists/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name     string
		l1       *golang.ListNode
		l2       *golang.ListNode
		expected *golang.ListNode
	}{
		{
			name:     "Example 1",
			l1:       &golang.ListNode{Val: 1, Next: &golang.ListNode{Val: 2, Next: &golang.ListNode{Val: 4}}},
			l2:       &golang.ListNode{Val: 1, Next: &golang.ListNode{Val: 3, Next: &golang.ListNode{Val: 4}}},
			expected: &golang.ListNode{Val: 1, Next: &golang.ListNode{Val: 1, Next: &golang.ListNode{Val: 2, Next: &golang.ListNode{Val: 3, Next: &golang.ListNode{Val: 4, Next: &golang.ListNode{Val: 4}}}}}},
		},
		{
			name:     "Example 2",
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			name:     "Example 3",
			l1:       nil,
			l2:       &golang.ListNode{Val: 0},
			expected: &golang.ListNode{Val: 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := golang.MergeTwoLists(tc.l1, tc.l2)
			require.Equal(t, listNodeToSlice(tc.expected), listNodeToSlice(result))
		})
	}
}

func listNodeToSlice(head *golang.ListNode) []int {
	if head == nil {
		return nil
	}

	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
