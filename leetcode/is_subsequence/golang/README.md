# Is Subsequence - Go Solution

## Algorithm Strategy

This solution uses a **greedy single-pointer approach** with rune conversion for Unicode safety.

The algorithm:
1. Returns false if s is longer than t (impossible to be subsequence)
2. Returns true if s is empty string
3. Converts both strings to rune slices (for proper Unicode handling)
4. Initializes `pointer` to track position in s
5. Iterates through each rune in t:
   - If current rune matches sRune[pointer], increments pointer
   - If pointer reaches len(s), returns true immediately
6. If loop completes without matching all of s, returns false

The key insight is the greedy approach: we match characters from left to right in order, and once a character is matched, we never need to reconsider it.

## Time Complexity

**O(n + m)** where n is the length of t and m is the length of s.

We iterate through t once (O(n)) and convert both strings to runes (O(n + m)).

## Space Complexity

**O(n + m)** for the rune slices.

Both strings are converted to rune arrays, which requires O(n + m) space.
