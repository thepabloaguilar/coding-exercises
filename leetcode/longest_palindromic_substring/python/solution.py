class Solution:
    def longestPalindrome(self, s: str) -> str:
        initial = 0
        final = 0

        max_length = 0
        for idx in range(len(s)):
            left, right = idx, idx
            while left >= 0 and right < len(s) and s[left] == s[right]:
                if (right - left) > max_length:
                    max_length = right - left
                    initial, final = left, right

                left -= 1
                right += 1

            left, right = idx, idx + 1
            while left >= 0 and right < len(s) and s[left] == s[right]:
                if (right - left) > max_length:
                    max_length = right - left
                    initial, final = left, right

                left -= 1
                right += 1

        return s[initial:final+1]
