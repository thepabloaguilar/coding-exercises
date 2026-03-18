package golang

import "math"

func MaxProfit(prices []int) int {
	profit := 0

	buy := int(math.Pow(10, 5))
	sell := 0
	for _, price := range prices {
		if price < buy {
			buy = price
			sell = price
		} else if price > sell {
			profit = max(profit, price-buy)
			sell = price
		}
	}

	return profit
}
