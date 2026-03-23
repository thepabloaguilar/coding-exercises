# Jump Game - Go Solution

## Algorithm Strategy

This solution uses a **greedy backward iteration** approach.

The algorithm:
1. Initializes `goal` to the last index (the target we want to reach)
2. Iterates backward from the last index to 0
3. For each position:
   - Checks if this position can reach the current `goal`
   - If `idx + nums[idx] >= goal`, updates goal to current index
4. Returns true if and only if goal equals 0 (we can reach the end from start)

The key insight is the greedy observation: if we work backwards and find that position i can reach the current goal, then position i becomes the new goal. We only need to check if we can eventually reach index 0.

## Time Complexity

**O(n)** where n is the length of the nums array.

Single backward iteration through the array.

## Space Complexity

**O(1)** constant space.

Only one variable (`goal`) stores state.
