class Solution:
    def hIndex(self, citations: list[int]) -> int:
        citations.sort(reverse=True)

        h = 0
        for idx, citation in enumerate(citations, start=1):
            if citation < idx:
                return h
            h = idx

        return h
