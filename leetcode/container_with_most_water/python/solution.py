class Solution:
    def maxArea(self, height: list[int]) -> int:
        left, right = 0, len(height) - 1
        maxContainer = min(height[left], height[right]) * (right - left)

        while left < right:
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1

            maxContainer = max(
                maxContainer,
                min(height[left], height[right]) * (right - left),
                )

        return maxContainer
