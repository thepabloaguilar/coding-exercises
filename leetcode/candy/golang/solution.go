package golang

func Candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for idx := 0; idx < len(ratings); idx++ {
		// Initializes candies to have at least one
		candies[idx] = 1

		if idx > 0 && ratings[idx] > ratings[idx-1] {
			candies[idx] = candies[idx-1] + 1
		}
	}

	candiesSum := 0
	for idx := len(ratings) - 1; idx >= 0; idx-- {
		if idx < len(ratings)-1 && ratings[idx] > ratings[idx+1] {
			candies[idx] = max(candies[idx], candies[idx+1]+1)
		}

		candiesSum += candies[idx]
	}

	return candiesSum
}
