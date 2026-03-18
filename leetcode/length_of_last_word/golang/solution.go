package golang

import "unicode"

func LengthOfLastWord(s string) int {
	start := 0
	end := 0

	sRune := []rune(s)
	for idx := len(sRune) - 1; idx >= 0; idx-- {
		if !unicode.IsSpace(sRune[idx]) && end == 0 {
			end = idx
		} else if end != 0 && unicode.IsSpace(sRune[idx]) {
			break
		}

		start = idx
	}

	return (end - start) + 1
}
