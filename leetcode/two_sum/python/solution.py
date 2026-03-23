class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        seenNums = dict()
        for idx, num in enumerate(nums):
            rest = target - num
            if rest in seenNums:
                return [seenNums[rest], idx]
            seenNums[num] = idx

        return [-1, -1]
