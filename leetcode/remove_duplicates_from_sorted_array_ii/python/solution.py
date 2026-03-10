class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        insertion_idx = 1
        last_value = nums[0]
        has_seen_the_same = False

        for search_idx in range(1, len(nums)):
            if nums[search_idx] != last_value:
                nums[insertion_idx] = nums[search_idx]
                last_value = nums[search_idx]
                insertion_idx += 1
                has_seen_the_same = False
            elif not has_seen_the_same:
                nums[insertion_idx] = last_value
                insertion_idx += 1
                has_seen_the_same = True

        return insertion_idx
