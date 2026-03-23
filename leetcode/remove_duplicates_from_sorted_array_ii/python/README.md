# Remove Duplicates from Sorted Array II - Python Solution

## Algorithm Strategy

This solution uses **two-pointer with duplicate tracking** allowing up to 2 occurrences.

The algorithm:
1. Initializes `insertion_idx` at 1, tracks `last_value` and `has_seen_the_same` flag
2. For each element from index 1:
   - If value differs from last_value:
     - Places it at insertion_idx, updates last_value
     - Increments insertion_idx, resets flag to False
   - Else if not yet seen duplicate (first duplicate):
     - Places it at insertion_idx, increments insertion_idx
     - Sets flag to True (second occurrence recorded)
   - Else (third+ occurrence): skips it
3. Returns insertion_idx

The key insight is tracking whether we've seen one duplicate already. The first two occurrences are kept, subsequent duplicates are skipped.

## Time Complexity

**O(n)** where n is the array length.

Single pass with constant-time operations per element.

## Space Complexity

**O(1)** constant space.

In-place modification using only a few variables.
