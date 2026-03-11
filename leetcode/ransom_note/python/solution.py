from collections import defaultdict

class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        if len(ransomNote) > len(magazine):
            return False

        counter = defaultdict(int)
        for letter in magazine:
            counter[letter] += 1

        for letter in ransomNote:
            counter[letter] -= 1
            if counter[letter] < 0:
                return False

        return True
