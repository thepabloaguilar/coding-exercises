# H-Index - Go Solution

## Algorithm Strategy

This solution uses a **bucket sort (counting sort) approach** for optimal time complexity.

The algorithm:
1. Creates a `buckets` array of size n+1 where n is the number of papers
2. For each citation:
   - Places it in bucket[min(citation, n)]
   - Citations ≥ n go in the last bucket
3. Iterates backward from h = n to 0:
   - Maintains `papersSum` (count of papers with ≥ h citations)
   - Increments papersSum by buckets[h]
   - When h ≤ papersSum, we've found the h-index
4. Returns the h-index

The key insight: The h-index definition means we need h papers with ≥ h citations. By counting papers in buckets and working backwards, we find the largest h satisfying this condition.

## Time Complexity

**O(n)** where n is the number of papers.

The algorithm makes two passes: one to fill buckets (O(n)) and one to find h-index (O(n)).

## Space Complexity

**O(n)** for the buckets array.

An additional array of size n+1 is used for counting.
