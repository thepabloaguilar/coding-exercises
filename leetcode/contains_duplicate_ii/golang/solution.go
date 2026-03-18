package golang

func ContainsNearbyDuplicate(nums []int, k int) bool {
	seenNumbers := make(map[int]int)
	for idx, num := range nums {
		seenIdx, ok := seenNumbers[num]
		if ok && idx-seenIdx <= k {
			return true
		}

		seenNumbers[num] = idx
	}

	return false
}
