# Length of Last Word - Go Solution

## Algorithm Strategy

This solution uses **backward iteration** to find the last word's boundaries.

The algorithm:
1. Converts string to rune slice for Unicode support
2. Iterates backward from the last character
3. Tracks two pointers:
   - `end`: Set when we find the first non-space character from the right
   - `start`: Continues moving left until we hit a space (after end is set)
4. When we encounter a space after finding the word's end, breaks
5. Returns `(end - start) + 1` as the word length

The key insight is scanning from right to left: skip trailing spaces, find the last word's end, then continue until finding its start.

## Time Complexity

**O(n)** where n is the length of the string.

In the worst case (entire string is one word), we scan the entire string once.

## Space Complexity

**O(n)** for the rune slice.

Converting the string to a rune slice requires O(n) space for proper Unicode handling.
