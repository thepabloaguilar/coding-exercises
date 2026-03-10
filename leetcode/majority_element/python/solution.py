from collections import defaultdict

class Solution:
    def majorityElement(self, nums: list[int]) -> int:
        threshold = len(nums) // 2
        counters = defaultdict(int)

        for num in nums:
            counters[num] += 1
            if counters[num] > threshold:
                return num

        raise Exception("No majority element")
