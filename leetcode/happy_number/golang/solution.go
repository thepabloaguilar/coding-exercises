package golang

import "math"

func IsHappy(n int) bool {
	seenNumbers := map[int]struct{}{
		n: {},
	}

	for n != 1 {
		n = sumOfSquares(n)
		if _, ok := seenNumbers[n]; ok {
			return false
		}
		seenNumbers[n] = struct{}{}
	}

	return true
}

func sumOfSquares(n int) int {
	result := 0
	for n > 0 {
		result += int(math.Pow(float64(n%10), 2))
		n /= 10
	}

	return result
}
