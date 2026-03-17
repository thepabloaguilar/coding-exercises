package golang

import "unicode"

func IsPalindrome(s string) bool {
	runes := []rune(s)
	start, end := 0, len(runes)-1
	for start < end {
		if !isAlphaNumeric(runes[start]) {
			start++
			continue
		}

		if !isAlphaNumeric(runes[end]) {
			end--
			continue
		}

		if unicode.ToLower(runes[start]) != unicode.ToLower(runes[end]) {
			return false
		}

		start++
		end--
	}

	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
