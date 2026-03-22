package golang

func Jump(nums []int) int {
	jumps := 0
	left, right := 0, 0
	for right < len(nums)-1 {
		farthest := 0
		for idx := left; idx <= right; idx++ {
			farthest = max(farthest, idx+nums[idx])
		}

		left = right + 1
		right = farthest
		jumps++
	}

	return jumps
}
