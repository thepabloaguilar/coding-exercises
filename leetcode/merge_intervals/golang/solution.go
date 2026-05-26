package golang

import (
	"cmp"
	"slices"
)

func Merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	result := [][]int{intervals[0]}
	for _, item := range intervals[1:] {
		lastEnd := result[len(result)-1][1]
		if item[0] <= lastEnd {
			result[len(result)-1][1] = max(lastEnd, item[1])
		} else {
			result = append(result, item)
		}
	}

	return result
}
