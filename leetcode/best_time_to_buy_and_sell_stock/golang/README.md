# Best Time to Buy and Sell Stock - Go Solution

## Algorithm Strategy

This solution uses a **single-pass greedy approach** with tracking of minimum buy price and maximum profit.

The algorithm:
1. Initializes `profit` to 0, `buy` to a large value (10^5), and `sell` to the initial price
2. Iterates through each price in the array
3. When a price lower than the current `buy` is found:
   - Updates `buy` to this new minimum
   - Resets `sell` to the same price (new buying point)
4. When a price higher than current `sell` is found:
   - Calculates potential profit (price - buy)
   - Updates `profit` to the maximum of current profit and new potential
   - Updates `sell` to the current price
5. Returns the maximum profit found

The key insight is to maintain the lowest purchase price encountered and calculate the maximum profit achievable by selling at any subsequent price.

## Time Complexity

**O(n)** where n is the length of the prices array.

The algorithm performs a single iteration through the prices array with constant-time operations for each element.

## Space Complexity

**O(1)** constant space.

The solution uses only three variables (`profit`, `buy`, `sell`) regardless of the input size.
