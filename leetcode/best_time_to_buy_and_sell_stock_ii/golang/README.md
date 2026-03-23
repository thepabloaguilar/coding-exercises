# Best Time to Buy and Sell Stock II - Go Solution

## Algorithm Strategy

This solution uses a **greedy approach** that accumulates all profitable price increases.

The algorithm:
1. Initializes `result` to 0 to track total profit
2. Iterates through indices from 0 to len(prices)-2
3. For each consecutive pair of days (idx, idx+1):
   - Checks if the next day's price is higher than today's
   - If yes, adds the price difference to the result
4. Returns the accumulated profit

The key insight is that maximum profit with unlimited transactions is achieved by summing all positive price differences between consecutive days. This effectively captures every upward price movement.

## Time Complexity

**O(n)** where n is the length of the prices array.

The algorithm performs a single iteration through the array, examining each consecutive pair once.

## Space Complexity

**O(1)** constant space.

The solution uses only a single `result` variable and loop index, independent of input size.
