# Valid Palindrome - Go Solution

## Algorithm Strategy

This solution uses **two-pointer with character filtering** for palindrome validation.

The algorithm:
1. Converts string to rune slice for Unicode support
2. Initializes start and end pointers
3. While start < end:
   - Skips non-alphanumeric runes from start
   - Skips non-alphanumeric runes from end
   - Compares lowercase versions using unicode.ToLower
   - If mismatch: returns false
   - Moves pointers inward
4. Returns true if all valid characters match

The isAlphaNumeric helper checks if a rune is letter or digit.

The key insight is on-the-fly filtering: skip invalid characters during comparison rather than preprocessing, saving space.

## Time Complexity

**O(n)** where n is the string length.

Each character examined once as pointers converge.

## Space Complexity

**O(n)** for the rune slice.

String to rune conversion requires O(n) space for Unicode handling.
