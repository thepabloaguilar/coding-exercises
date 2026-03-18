class Solution:
    def containsNearbyDuplicate(self, nums: list[int], k: int) -> bool:
        seen_numbers = dict()
        for idx, num in enumerate(nums):
            seen_idx = seen_numbers.get(num)
            if seen_idx is not None and abs(seen_idx - idx) <= k:
                return True
            else:
                seen_numbers[num] = idx

        return False
