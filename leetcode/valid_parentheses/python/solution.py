from collections import deque


class Solution:
    def isValid(self, s: str) -> bool:
        stack =  deque()

        pairs = {
            ')': '(',
            ']': '[',
            '}': '{',
        }
        closing = set(pairs.keys())
        for letter in s:
            if letter in closing:
                if not stack or stack[-1] != pairs[letter]:
                    return False
                stack.pop()
            else:
                stack.append(letter)

        return len(stack) == 0
