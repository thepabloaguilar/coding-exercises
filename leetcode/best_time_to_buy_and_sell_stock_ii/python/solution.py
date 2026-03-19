class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        result = 0
        for idx in range(len(prices) - 1):
            if prices[idx] < prices[idx+1]:
                result += prices[idx+1] - prices[idx]

        return result
