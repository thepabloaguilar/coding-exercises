package golang

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for haystack_idx := 0; haystack_idx < len(haystack)-len(needle)+1; haystack_idx++ {
		for needle_idx := 0; needle_idx < len(needle); needle_idx++ {
			if haystack[haystack_idx+needle_idx] != needle[needle_idx] {
				break
			}

			if needle_idx == len(needle)-1 {
				return haystack_idx
			}
		}
	}

	return -1
}
