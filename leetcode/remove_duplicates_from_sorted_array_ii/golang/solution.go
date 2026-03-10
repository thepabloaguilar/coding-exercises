package golang

func RemoveDuplicates(nums []int) int {
	insertionIdx := 0

	for _, num := range nums {
		if insertionIdx < 2 || num != nums[insertionIdx-2] {
			nums[insertionIdx] = num
			insertionIdx++
		}
	}

	return insertionIdx
}
