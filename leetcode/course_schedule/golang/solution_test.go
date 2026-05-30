package golang_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepabloaguilar/coding-exercises/leetcode/course_schedule/golang"
)

func TestBaseCases(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:       "Example 1",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
			expected: true,
		},
		{
			name:       "Example 2",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.expected, golang.CanFinish(tc.numCourses, tc.prerequisites))
		})
	}
}
