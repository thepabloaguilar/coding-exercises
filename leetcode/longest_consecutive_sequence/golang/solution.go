package golang

import "slices"

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)

	maxLength := 1
	currentLength := 1
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == nums[idx-1] {
			continue
		} else if nums[idx] == nums[idx-1]+1 {
			currentLength++
			maxLength = max(maxLength, currentLength)
		} else {
			currentLength = 1
		}
	}

	return maxLength
}
