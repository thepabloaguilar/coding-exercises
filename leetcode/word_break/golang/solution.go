package golang

func WordBreak(s string, wordDict []string) bool {
	matched := make([]bool, len(s)+1)
	matched[len(s)] = true

	for idx := len(s) - 1; idx >= 0; idx-- {
		for _, word := range wordDict {
			if idx+len(word) <= len(s) && s[idx:idx+len(word)] == word {
				matched[idx] = matched[idx+len(word)]
			}

			if matched[idx] {
				break
			}
		}
	}

	return matched[0]
}
