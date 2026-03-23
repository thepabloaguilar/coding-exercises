# Jump Game - Python Solution

## Algorithm Strategy

This solution uses a **greedy backward iteration** approach.

The algorithm:
1. Initializes `goal` to the last index (target position)
2. Iterates backward from the second-to-last index to 0
3. For each position:
   - Checks if we can reach the current `goal` from this position
   - If `idx + nums[idx] >= goal`, updates goal to this index
4. Returns True if goal reaches index 0 (we can reach the end from the start)

The key insight is working backwards: if we can reach position i, and from position j we can reach position i, then we only need to check if we can reach position j. This greedy approach simplifies the problem.

## Time Complexity

**O(n)** where n is the length of the nums array.

Single backward pass through the array with constant-time operations.

## Space Complexity

**O(1)** constant space.

Only one variable (`goal`) is used regardless of input size.
