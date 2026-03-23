# Best Time to Buy and Sell Stock II - Python Solution

## Algorithm Strategy

This solution uses a **greedy approach** that captures every profitable upward price movement.

The algorithm:
1. Initializes `result` to 0 to accumulate total profit
2. Iterates through the prices array from index 0 to n-2
3. For each consecutive pair of days (idx, idx+1):
   - If tomorrow's price is higher than today's price
   - Adds the difference (profit) to the result
4. Returns the accumulated profit

The key insight is that to maximize profit with unlimited transactions, we should capture every single upward price movement. This is equivalent to buying before every price increase and selling right before it drops.

## Time Complexity

**O(n)** where n is the length of the prices array.

The algorithm makes a single pass through the array, comparing each adjacent pair of prices exactly once.

## Space Complexity

**O(1)** constant space.

The solution only uses a single variable `result` and the loop index, regardless of input size.
