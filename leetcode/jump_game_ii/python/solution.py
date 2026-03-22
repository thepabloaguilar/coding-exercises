class Solution:
    def jump(self, nums: list[int]) -> int:
        jumps = 0
        left = right = 0
        while right < len(nums)-1:
            farthest = 0
            for idx in range(left, right + 1):
                farthest = max(farthest, idx + nums[idx])

            left = right + 1
            right = farthest
            jumps += 1

        return jumps
