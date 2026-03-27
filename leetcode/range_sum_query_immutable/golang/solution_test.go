package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/range_sum_query_immutable/golang"
)

type query struct {
	left     int
	right    int
	expected int
}

func TestBaseCases(t *testing.T) {
	testCases := []struct {
		name  string
		nums  []int
		query []query
	}{
		{
			name: "Example 1",
			nums: []int{-2, 0, 3, -5, 2, -1},
			query: []query{
				{left: 0, right: 2, expected: 1},
				{left: 2, right: 5, expected: -1},
				{left: 0, right: 5, expected: -3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			numArray := golang.Constructor(tc.nums)
			for _, q := range tc.query {
				require.Equal(t, q.expected, numArray.SumRange(q.left, q.right))
			}
		})
	}
}
