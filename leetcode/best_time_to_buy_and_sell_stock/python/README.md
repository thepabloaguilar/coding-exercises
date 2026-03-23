# Best Time to Buy and Sell Stock - Python Solution

## Algorithm Strategy

This solution uses a **single-pass greedy approach** with tracking of minimum buy price and maximum profit.

The algorithm:
1. Initializes `profit` to 0, `buy` to a large value (10^5), and `sell` to 0
2. Iterates through each price in the array
3. When a lower price is found:
   - Updates the `buy` price to this new minimum
   - Resets `sell` to 0 (potential new buying opportunity)
4. When a price higher than current `sell` is found:
   - Calculates potential profit (price - buy)
   - Updates `profit` to the maximum of current profit and new potential
   - Updates `sell` to the current price
5. Returns the maximum profit found

The key insight is to always track the minimum purchase price seen so far and calculate the profit if we were to sell at each subsequent price.

## Time Complexity

**O(n)** where n is the length of the prices array.

The algorithm makes a single pass through the prices array, performing constant-time operations for each element.

## Space Complexity

**O(1)** constant space.

The solution only uses three variables (`profit`, `buy`, `sell`) regardless of input size.
