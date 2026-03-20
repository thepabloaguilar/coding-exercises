package golang

func HIndex(citations []int) int {
	buckets := make([]int, len(citations)+1)
	for _, citation := range citations {
		bucketIdx := min(citation, len(citations))
		buckets[bucketIdx]++
	}

	h := len(citations)
	papersSum := buckets[h]
	for h > papersSum {
		h--
		papersSum += buckets[h]
	}

	return h
}
