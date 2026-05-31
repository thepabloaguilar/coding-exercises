from collections import defaultdict


class Solution:
    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        counter = defaultdict(int)
        for num in nums:
            counter[num] += 1

        buckets = [[] for _ in range(len(nums) + 1)]
        for number, frequency in counter.items():
            buckets[frequency].append(number)

        result = []
        for idx in range(len(buckets) - 1, 0, -1):
            for item in buckets[idx]:
                result.append(item)
                if len(result) == k:
                    return result

        return result
