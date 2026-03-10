class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        insertion_idx = 0
        search_idx = 0

        while search_idx < len(nums):
            if nums[search_idx] != val:
                nums[insertion_idx] = nums[search_idx]
                insertion_idx += 1
            search_idx += 1

        return insertion_idx
