package golang

func MaxProfit(prices []int) int {
	result := 0
	for idx := range len(prices) - 1 {
		if prices[idx] < prices[idx+1] {
			result += prices[idx+1] - prices[idx]
		}
	}

	return result
}
