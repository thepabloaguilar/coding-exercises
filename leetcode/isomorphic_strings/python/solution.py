class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        mapping = {}
        seen_letters = set()
        for idx in range(len(s)):
            letter = s[idx]

            mapping_letter = mapping.get(letter, None)
            if not mapping_letter:
                if t[idx] in seen_letters:
                    return False

                mapping[letter] = t[idx]
                seen_letters.add(t[idx])
            elif mapping_letter != t[idx]:
                return False

        return True
