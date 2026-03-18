class Solution:
    def maxProfit(self, prices: list[int]) -> int:
        profit = 0

        buy = 10 ** 5
        sell = 0
        for price in prices:
            if price < buy:
                buy = price
                sell = 0

            if price > sell:
                profit = max(profit, price - buy)
                sell = price

        return profit
