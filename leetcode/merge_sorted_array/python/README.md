# Merge Sorted Array - Python Solution

## Algorithm Strategy

This solution uses **backward three-pointer merging** to merge in-place.

The algorithm:
1. Initializes `idx` to the last position of the merged array (m + n - 1)
2. Adjusts m and n to be 0-indexed (decrements both)
3. While idx >= 0:
   - If m exhausted but n has elements: places nums2[n] and decrements n
   - If n exhausted but m has elements: places nums1[m] and decrements m
   - If both have elements: compares nums1[m] vs nums2[n]
     - Places the larger value and decrements corresponding pointer
   - Decrements idx
4. Modifies nums1 in-place

The key insight is working backwards from the end to avoid overwriting unprocessed elements in nums1. We always have space at the end to place the larger element.

## Time Complexity

**O(m + n)** where m and n are the lengths of the two arrays.

We process each element from both arrays exactly once.

## Space Complexity

**O(1)** constant space.

The merge is done in-place using only a few pointer variables.
