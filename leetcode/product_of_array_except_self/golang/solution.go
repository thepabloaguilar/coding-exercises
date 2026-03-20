package golang

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for idx := 0; idx < len(nums); idx++ {
		result[idx] = prefix
		prefix *= nums[idx]
	}

	postfix := 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		result[idx] *= postfix
		postfix *= nums[idx]
	}

	return result
}
