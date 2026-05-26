class Solution:
    def merge(self, intervals: list[list[int]]) -> list[list[int]]:
        intervals.sort(key=lambda item: item[0])

        result = [intervals[0]]
        for currentStart, currentEnd in intervals[1:]:
            lastEnd = result[-1][1]
            if currentStart <= lastEnd:
                result[-1][1] = max(currentEnd, lastEnd)
            else:
                result.append([currentStart, currentEnd])

        return result
