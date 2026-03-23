# Jump Game II - Go Solution

## Algorithm Strategy

This solution uses a **greedy BFS-style approach** to minimize jumps.

The algorithm:
1. Initializes `jumps` to 0 and a window `[left, right]` at [0, 0]
2. While right is less than the last index:
   - Finds the `farthest` position reachable from the current window
   - Iterates through all positions in [left, right]
   - For each position, calculates the maximum of `farthest` and `idx + nums[idx]`
   - Shifts the window to the next level: `left = right + 1`, `right = farthest`
   - Increments `jumps` counter
3. Returns the total jumps needed

The key insight is level-by-level exploration: all positions within the current window [left, right] are reachable with the same number of jumps. We greedily expand to the farthest reachable position in each jump.

## Time Complexity

**O(n)** where n is the length of the nums array.

Despite nested loops, each element is processed at most once as windows partition the array.

## Space Complexity

**O(1)** constant space.

Only three variables (jumps, left, right, farthest) are used.
