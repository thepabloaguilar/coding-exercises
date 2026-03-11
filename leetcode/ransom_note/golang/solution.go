package golang

func CanConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	counters := make(map[rune]int)

	for _, letter := range magazine {
		counters[letter]++
	}

	for _, letter := range ransomNote {
		counters[letter]--
		if counters[letter] < 0 {
			return false
		}
	}

	return true
}
