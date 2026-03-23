class Solution:
    def trap(self, height: list[int]) -> int:
        max_left = [0] * len(height)
        last_max_left = 0
        for idx in range(1, len(height)):
            last_max_left = max(last_max_left, height[idx-1])
            max_left[idx] = last_max_left

        max_right = [0] * len(height)
        last_max_right = 0
        for idx in range(len(height) - 2, -1, -1):
            last_max_right = max(last_max_right, height[idx+1])
            max_right[idx] = last_max_right

        result = 0
        for idx in range(len(height)):
            water = min(max_left[idx], max_right[idx]) - height[idx]
            if water > 0:
                result += water

        return result
