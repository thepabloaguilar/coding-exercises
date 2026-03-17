package golang

func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if s == "" {
		return true
	}

	pointer := 0
	sRune := []rune(s)
	tRune := []rune(t)
	for _, letter := range tRune {
		if sRune[pointer] == letter {
			pointer += 1

			if pointer >= len(s) {
				return true
			}
		}
	}

	return false
}
