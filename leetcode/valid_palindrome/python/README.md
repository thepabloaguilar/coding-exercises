# Valid Palindrome - Python Solution

## Algorithm Strategy

This solution uses **two-pointer convergence** with character filtering.

The algorithm:
1. Initializes start and end pointers at both ends
2. While start < end:
   - Skips non-alphanumeric characters from start (increments)
   - Skips non-alphanumeric characters from end (decrements)
   - Compares lowercase versions of characters
   - If mismatch: returns False
   - Moves both pointers inward
3. Returns True if all comparisons match

The key insight is filtering: skip non-alphanumeric characters on the fly using isalnum(), compare case-insensitively using lower().

## Time Complexity

**O(n)** where n is the string length.

Each character examined at most once as pointers converge.

## Space Complexity

**O(1)** constant space.

Only pointer variables, no additional data structures.
