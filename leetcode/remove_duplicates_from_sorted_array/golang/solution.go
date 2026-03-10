package golang

func RemoveDuplicates(nums []int) int {
	insertionIdx := 1
	searchIdx := 1
	lastValue := nums[0]

	for searchIdx < len(nums) {
		if nums[searchIdx] != lastValue {
			nums[insertionIdx] = nums[searchIdx]
			lastValue = nums[searchIdx]
			insertionIdx++
		}

		searchIdx++
	}

	return insertionIdx
}
