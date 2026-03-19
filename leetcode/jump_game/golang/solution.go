package golang

func CanJump(nums []int) bool {
	goal := len(nums) - 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		if idx+nums[idx] >= goal {
			goal = idx
		}
	}

	return goal == 0
}
