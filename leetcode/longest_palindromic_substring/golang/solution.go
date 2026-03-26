package golang

func LongestPalindrome(s string) string {
	initial, final := 0, 0
	maxLength := 0

	for idx := range s {
		left, right := idx, idx
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left) > maxLength {
				initial, final = left, right
				maxLength = right - left
			}

			left--
			right++
		}

		left, right = idx, idx+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left) > maxLength {
				initial, final = left, right
				maxLength = right - left
			}

			left--
			right++
		}
	}

	return s[initial : final+1]
}
