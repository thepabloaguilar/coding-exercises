package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/add_two_numbers/golang"
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
			l1:       &golang.ListNode{Val: 2, Next: &golang.ListNode{Val: 4, Next: &golang.ListNode{Val: 3}}},
			l2:       &golang.ListNode{Val: 5, Next: &golang.ListNode{Val: 6, Next: &golang.ListNode{Val: 4}}},
			expected: &golang.ListNode{Val: 7, Next: &golang.ListNode{Val: 0, Next: &golang.ListNode{Val: 8}}},
		},
		{
			name:     "Example 2",
			l1:       &golang.ListNode{Val: 0},
			l2:       &golang.ListNode{Val: 0},
			expected: &golang.ListNode{Val: 0},
		},
		{
			name: "Example 3",
			l1: &golang.ListNode{
				Val: 9,
				Next: &golang.ListNode{
					Val: 9,
					Next: &golang.ListNode{
						Val: 9,
						Next: &golang.ListNode{
							Val: 9,
							Next: &golang.ListNode{
								Val: 9,
								Next: &golang.ListNode{
									Val:  9,
									Next: &golang.ListNode{Val: 9},
								},
							},
						},
					},
				},
			},
			l2: &golang.ListNode{
				Val: 9,
				Next: &golang.ListNode{
					Val: 9,
					Next: &golang.ListNode{
						Val:  9,
						Next: &golang.ListNode{Val: 9},
					},
				},
			},
			expected: &golang.ListNode{
				Val: 8,
				Next: &golang.ListNode{
					Val: 9,
					Next: &golang.ListNode{
						Val: 9,
						Next: &golang.ListNode{
							Val: 9,
							Next: &golang.ListNode{
								Val: 0,
								Next: &golang.ListNode{
									Val: 0,
									Next: &golang.ListNode{
										Val:  0,
										Next: &golang.ListNode{Val: 1},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := golang.AddTwoNumbers(tc.l1, tc.l2)
			require.Equal(t, listNodeToSlice(tc.expected), listNodeToSlice(result))
		})
	}
}

func listNodeToSlice(head *golang.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
