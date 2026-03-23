# Remove Element - Python Solution

## Algorithm Strategy

This solution uses a **two-pointer filtering approach** for in-place removal.

The algorithm:
1. Initializes both `insertion_idx` and `search_idx` at 0
2. While search_idx < len(nums):
   - If nums[search_idx] != val:
     - Copies it to insertion_idx position
     - Increments insertion_idx
   - Always increments search_idx
3. Returns insertion_idx (count of remaining elements)

The key insight is maintaining two pointers: search_idx scans all elements, insertion_idx tracks where to place the next non-val element. Elements equal to val are simply skipped.

## Time Complexity

**O(n)** where n is the array length.

Single pass through the array.

## Space Complexity

**O(1)** constant space.

In-place modification with only pointer variables.
