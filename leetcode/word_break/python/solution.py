class Solution:
    def wordBreak(self, s: str, wordDict: list[str]) -> bool:
        matched = [False] * (len(s) + 1)
        matched[-1] = True

        for idx in range(len(s) - 1, -1, -1):
            for word in wordDict:
                if idx + len(word) <= len(s) and s[idx:idx+len(word)] == word:
                    matched[idx] = matched[idx + len(word)]

                if matched[idx]:
                    break

        return matched[0]
