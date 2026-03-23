package golang

func TwoSum(nums []int, target int) []int {
	seenNums := make(map[int]int)
	for idx, num := range nums {
		rest := target - num
		if restIdx, ok := seenNums[rest]; ok {
			return []int{restIdx, idx}
		}
		seenNums[num] = idx
	}

	return []int{-1, -1}
}
