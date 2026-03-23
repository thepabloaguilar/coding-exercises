# Remove Element - Go Solution

## Algorithm Strategy

This solution uses a **swap-with-end approach** for in-place removal.

The algorithm:
1. Initializes `itemToChange` at last index
2. Iterates i from 0 to itemToChange:
   - If nums[i] == val:
     - Swaps with element at itemToChange
     - Decrements itemToChange
     - Decrements i (to recheck swapped element)
   - Otherwise continues forward
3. Returns itemToChange + 1 (remaining element count)

The key insight is swapping unwanted elements to the end: when we find val, swap it with the last unprocessed element and shrink the active range. The decrement i ensures we check the swapped value.

## Time Complexity

**O(n)** where n is the array length.

Each element is processed at most once (some may be swapped but not rechecked if they're also val).

## Space Complexity

**O(1)** constant space.

In-place swapping with only pointer variables.
