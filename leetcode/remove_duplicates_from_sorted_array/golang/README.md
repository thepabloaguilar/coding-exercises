# Remove Duplicates from Sorted Array - Go Solution

## Algorithm Strategy

This solution uses **two-pointer in-place array modification**.

The algorithm:
1. Initializes `insertionIdx` at 1, `searchIdx` at 1, and `lastValue` to nums[0]
2. While searchIdx < len(nums):
   - If nums[searchIdx] differs from lastValue:
     - Places it at insertionIdx
     - Updates lastValue
     - Increments insertionIdx
   - Always increments searchIdx
3. Returns insertionIdx (number of unique elements)

The key insight is that in a sorted array, duplicates are consecutive. By comparing each element with the last unique value, we can filter duplicates in one pass.

## Time Complexity

**O(n)** where n is the array length.

Single iteration through the array.

## Space Complexity

**O(1)** constant space.

In-place modification with only pointer variables.
