# Gas Station - Python Solution

## Algorithm Strategy

This solution uses a **greedy single-pass approach** to find the starting gas station.

The algorithm:
1. First checks if a solution exists: if sum(gas) < sum(cost), returns -1
2. Initializes `start_idx` to 0 and `total` (tank) to 0
3. Iterates through each station:
   - Calculates net gas: gas[idx] - cost[idx]
   - Adds net gas to total tank
   - If total becomes negative:
     - Current starting point cannot reach this station
     - Resets total to 0
     - Sets new starting point to idx + 1
4. Returns the final `start_idx`

The key insight is that if we run out of gas at station i when starting from station j, then any station between j and i also cannot be a valid starting point. Therefore, we should start fresh from i+1.

## Time Complexity

**O(n)** where n is the number of gas stations.

The algorithm makes two passes: one to calculate sums (implicit in the sum() calls) and one main iteration through all stations.

## Space Complexity

**O(1)** constant space.

Only a few variables are used regardless of input size.
