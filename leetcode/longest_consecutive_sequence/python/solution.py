class Solution:
    def longestConsecutive(self, nums: list[int]) -> int:
        seenNums = set(nums)
        result = 0

        for num in seenNums:
            if num - 1 not in seenNums:
                length = 1
                while (num + length) in seenNums:
                    length += 1
                result = max(result, length)

        return result
