# Roman to Integer - Go Solution

## Algorithm Strategy

This solution uses **lookahead comparison** to handle subtractive notation.

The algorithm:
1. Creates valueMapping for Roman numeral characters
2. Converts string to rune slice
3. For each character (rune) with index:
   - If not last character AND current value < next value:
     - Subtracts current value (subtractive pair like IV)
   - Otherwise: adds current value
4. Returns accumulated result

The key insight is lookahead: by checking if current < next, we detect subtractive notation (I before V means 4, not 6). Subtract the smaller value when it precedes a larger value, otherwise add.

## Time Complexity

**O(n)** where n is the string length.

Single pass with constant-time map lookups and comparisons.

## Space Complexity

**O(n)** for the rune slice conversion.

Converting string to rune slice requires O(n) space for Unicode handling.
