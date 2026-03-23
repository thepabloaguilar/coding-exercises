# Remove Duplicates from Sorted Array II - Go Solution

## Algorithm Strategy

This solution uses an **elegant two-pointer approach** comparing elements two positions apart.

The algorithm:
1. Initializes `insertionIdx` at 0
2. For each num in the array:
   - Condition: insertionIdx < 2 OR num != nums[insertionIdx-2]
   - If true: places num at insertionIdx and increments
   - This allows at most 2 consecutive duplicates
3. Returns insertionIdx (length of modified array)

The key insight: by comparing with the element 2 positions back, we ensure at most 2 consecutive equal values. If current differs from nums[insertionIdx-2], it's safe to add (at most 2 copies exist).

## Time Complexity

**O(n)** where n is the array length.

Single iteration through the array.

## Space Complexity

**O(1)** constant space.

In-place modification with only insertion index variable.
