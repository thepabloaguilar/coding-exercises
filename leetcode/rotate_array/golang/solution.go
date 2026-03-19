package golang

import "slices"

func Rotate(nums []int, k int) {
	k = k % len(nums)

	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
