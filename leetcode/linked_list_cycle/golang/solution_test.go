package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/linked_list_cycle/golang"
)

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name      string
		buildList func() *golang.ListNode
		expected  bool
	}{
		{
			name: "Example 1",
			buildList: func() *golang.ListNode {
				head := &golang.ListNode{Val: 3}
				head.Next = &golang.ListNode{Val: 2}
				head.Next.Next = &golang.ListNode{Val: 0}
				head.Next.Next.Next = &golang.ListNode{Val: -4}

				// Link node 1 to node 3
				head.Next.Next.Next.Next = head.Next

				return head
			},
			expected: true,
		},
		{
			name: "Example 2",
			buildList: func() *golang.ListNode {
				head := &golang.ListNode{Val: 1}
				head.Next = &golang.ListNode{Val: 2}

				// Link node 0 to node 1
				head.Next.Next = head

				return head
			},
			expected: true,
		},
		{
			name: "Example 3",
			buildList: func() *golang.ListNode {
				return &golang.ListNode{Val: 1}
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := golang.HasCycle(tc.buildList())
			require.Equal(t, tc.expected, result)
		})
	}
}
