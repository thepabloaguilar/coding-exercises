# Merge Sorted Array - Go Solution

## Algorithm Strategy

This solution uses **backward merging with three pointers** for in-place merging.

The algorithm:
1. Initializes `idx` at position m + n - 1 (last index of merged array)
2. Decrements m and n to convert to 0-indexed positions
3. While n >= 0 (while nums2 has unprocessed elements):
   - If m >= 0 and nums1[m] > nums2[n]:
     - Places nums1[m] at idx and decrements m
   - Otherwise:
     - Places nums2[n] at idx and decrements n
   - Decrements idx
4. Modifies nums1 in-place

The key insight is iterating backwards to avoid overwriting unprocessed elements. The condition focuses on placing all nums2 elements because any remaining nums1 elements are already in correct position.

## Time Complexity

**O(m + n)** where m and n are the array lengths.

Each element from both arrays is processed exactly once.

## Space Complexity

**O(1)** constant space.

In-place merging using only pointer variables.
