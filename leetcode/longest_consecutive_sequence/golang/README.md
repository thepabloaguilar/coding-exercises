# Longest Consecutive Sequence - Go Solution

## Algorithm Strategy

This solution uses a **sorting-based approach** to find consecutive sequences.

The algorithm:
1. Returns 0 if array is empty
2. Sorts the nums array in ascending order
3. Initializes `maxLength` and `currentLength` to 1
4. Iterates from index 1 to n-1:
   - If current equals previous (duplicate): skips it
   - If current equals previous + 1 (consecutive): increments currentLength and updates maxLength
   - Otherwise (break in sequence): resets currentLength to 1
5. Returns the maximum length found

The key insight is that after sorting, consecutive sequences appear as adjacent elements with values differing by 1.

## Time Complexity

**O(n log n)** where n is the length of the nums array.

Dominated by the sorting operation. The subsequent linear scan is O(n).

## Space Complexity

**O(1)** or **O(log n)** depending on the sort implementation.

The sort may use O(log n) stack space for quicksort. No additional data structures are used.
