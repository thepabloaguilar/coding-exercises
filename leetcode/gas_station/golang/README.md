# Gas Station - Go Solution

## Algorithm Strategy

This solution uses a **greedy single-pass approach** with running totals.

The algorithm:
1. Initializes `total` (overall balance), `tank` (current journey), and `startIdx` to 0
2. Iterates through each station:
   - Calculates diff: gas[idx] - cost[idx]
   - Updates both `total` and `tank` with diff
   - If tank becomes negative:
     - Current starting point fails
     - Resets tank to 0
     - Sets new start to idx + 1
3. After the loop, checks if total < 0:
   - If yes, no solution exists, returns -1
   - Otherwise, returns `startIdx`

The key insights:
1. If total gas ≥ total cost, a solution must exist
2. If we fail to reach station i from station j, any station between j and i also fails
3. Therefore, the next potential start is i+1

## Time Complexity

**O(n)** where n is the number of gas stations.

Single pass through all stations with constant-time operations at each step.

## Space Complexity

**O(1)** constant space.

Only three variables are used regardless of input size.
