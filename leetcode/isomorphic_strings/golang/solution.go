package golang

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRune := []rune(s)
	tRune := []rune(t)

	mappingS := make(map[rune]rune)
	mappingT := make(map[rune]rune)

	for idx := range len(s) {
		letter := sRune[idx]

		mappingLetter, ok := mappingS[letter]
		if !ok {
			if _, ok := mappingT[tRune[idx]]; ok {
				return false
			}

			mappingS[letter] = tRune[idx]
			mappingT[tRune[idx]] = letter
		} else if mappingLetter != tRune[idx] {
			return false
		}
	}

	return true
}
