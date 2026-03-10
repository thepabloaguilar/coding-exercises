class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        insertion_idx = 1
        search_idx = 1
        last_value = nums[0]

        while search_idx < len(nums):
            if nums[search_idx] != last_value:
                nums[insertion_idx] = nums[search_idx]
                last_value = nums[search_idx]
                insertion_idx += 1

            search_idx += 1

        return insertion_idx
