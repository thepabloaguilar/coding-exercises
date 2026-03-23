# Is Subsequence - Python Solution

## Algorithm Strategy

This solution uses a **greedy two-pointer approach** (single pointer with iteration).

The algorithm:
1. Returns False if s is longer than t (impossible to be subsequence)
2. Returns True if s is empty (empty string is subsequence of any string)
3. Initializes `pointer` to track position in s
4. Iterates through each character in t:
   - If current character matches s[pointer], increments pointer
   - If pointer reaches end of s, returns True (all characters matched)
5. If loop completes without matching all of s, returns False

The key insight is that we can greedily match characters from left to right. Once we match a character in s, we move to the next character and never backtrack.

## Time Complexity

**O(n)** where n is the length of string t.

We iterate through t at most once, performing constant-time operations for each character.

## Space Complexity

**O(1)** constant space.

Only a single pointer variable is used, regardless of input size.
