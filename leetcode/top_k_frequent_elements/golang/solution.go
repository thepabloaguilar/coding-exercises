package golang

func TopKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for number, frequency := range counter {
		buckets[frequency] = append(buckets[frequency], number)
	}

	result := make([]int, 0, k)
	for idx := len(buckets) - 1; idx > 0; idx-- {
		for _, num := range buckets[idx] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}

	return nil
}
