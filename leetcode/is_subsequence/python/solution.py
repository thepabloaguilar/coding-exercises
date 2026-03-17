class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) > len(t):
            return False

        if not s:
            return True

        pointer = 0
        for letter in t:
            if s[pointer] == letter:
                pointer += 1

            if pointer >= len(s):
                return True

        return False
