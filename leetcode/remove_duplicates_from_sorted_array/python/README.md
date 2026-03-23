# Remove Duplicates from Sorted Array - Python Solution

## Algorithm Strategy

This solution uses a **two-pointer in-place modification** approach.

The algorithm:
1. Initializes `insertion_idx` at 1 (position for next unique element)
2. Tracks `last_value` as nums[0]
3. Iterates with `search_idx` from 1 to n-1:
   - If current value differs from last_value:
     - Places it at insertion_idx
     - Updates last_value
     - Increments insertion_idx
   - Always increments search_idx
4. Returns insertion_idx (count of unique elements)

The key insight is exploiting sorted order: duplicates are adjacent, so we only keep the first occurrence of each value by comparing with the last unique value seen.

## Time Complexity

**O(n)** where n is the array length.

Single pass through the array with constant-time operations.

## Space Complexity

**O(1)** constant space.

In-place modification using only a few pointer variables.
