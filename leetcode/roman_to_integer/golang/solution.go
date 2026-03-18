package golang

func RomanToInt(s string) int {
	result := 0

	valueMapping := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sRune := []rune(s)
	for idx := range len(sRune) {
		if idx < len(sRune)-1 && valueMapping[sRune[idx]] < valueMapping[sRune[idx+1]] {
			result -= valueMapping[sRune[idx]]
		} else {
			result += valueMapping[sRune[idx]]
		}
	}

	return result
}
