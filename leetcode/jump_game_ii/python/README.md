# Jump Game II - Python Solution

## Algorithm Strategy

This solution uses a **greedy BFS-like approach** to find the minimum number of jumps.

The algorithm:
1. Initializes `jumps` to 0 and maintains a window `[left, right]` starting at [0, 0]
2. While right hasn't reached the last index:
   - Finds the farthest position reachable from any index in the current window
   - For each index in [left, right], calculates `idx + nums[idx]`
   - Updates `farthest` to the maximum reachable position
   - Moves the window: `left = right + 1`, `right = farthest`
   - Increments `jumps` (we made one jump to expand our reach)
3. Returns the total number of jumps

The key insight is treating this like BFS levels: all positions reachable in the current window can be reached with the same number of jumps. We greedily jump to the farthest point in the next level.

## Time Complexity

**O(n)** where n is the length of the nums array.

Although there are nested loops, each index is visited at most once as the windows don't overlap.

## Space Complexity

**O(1)** constant space.

Only a few variables are used regardless of input size.
