# H-Index - Python Solution

## Algorithm Strategy

This solution uses a **sorting-based approach** to calculate the h-index.

The algorithm:
1. Sorts the citations array in descending order (highest to lowest)
2. Initializes h-index `h` to 0
3. Iterates through citations with 1-based indexing (enumerate starts at 1)
4. For each citation:
   - If citation < idx (paper position):
     - We've found the point where papers have fewer citations than their position
     - Returns the current h-index value
   - Otherwise, updates h to idx (this paper qualifies for h-index of idx)
5. If all papers qualify, returns the final h value

The key insight: After sorting in descending order, the h-index is the largest value of `idx` where the citation count is ≥ idx.

## Time Complexity

**O(n log n)** where n is the number of papers.

Dominated by the sorting operation. The subsequent iteration is O(n).

## Space Complexity

**O(1)** or **O(n)** depending on the sorting algorithm.

Python's sort modifies the array in-place but may use O(n) temporary space for Timsort. The solution itself uses only constant additional space.
