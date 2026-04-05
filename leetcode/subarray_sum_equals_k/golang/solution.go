package golang

func SubarraySum(nums []int, k int) int {
	result := 0
	currentSum := 0
	prefixSums := map[int]int{0: 1}

	for _, num := range nums {
		currentSum += num
		diff := currentSum - k

		result += prefixSums[diff]
		prefixSums[currentSum]++
	}

	return result
}
