package golang

func IsValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0)

	for _, letter := range s {
		if letter == ')' || letter == ']' || letter == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[letter] {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, letter)
		}
	}

	return len(stack) == 0
}
