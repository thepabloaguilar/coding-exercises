# Happy Number - Python Solution

## Algorithm Strategy

This solution uses **cycle detection with a hash set** to determine if a number is happy.

The algorithm:
1. Initializes an empty set `seen_numbers` to track visited numbers
2. While n is not 1:
   - Calculates the sum of squares of digits using helper method `_sumOfSquares`
   - If this sum has been seen before:
     - We've detected a cycle, returns False
   - Adds the sum to seen_numbers
   - Updates n to this sum
3. If we reach 1, returns True

The helper method `_sumOfSquares`:
- Extracts each digit using modulo 10
- Squares it and adds to result
- Divides by 10 to remove the processed digit
- Continues until no digits remain

The key insight is that if we encounter the same number twice, we're in an infinite cycle and the number is not happy.

## Time Complexity

**O(log n)** where n is the input number.

The number of digits is O(log n). Each iteration reduces the magnitude significantly (sum of squares of digits is much smaller than the original number for large numbers). The cycle detection ensures we don't loop indefinitely.

## Space Complexity

**O(log n)** for the set storing seen numbers.

The set size is bounded by the number of unique values encountered before reaching 1 or detecting a cycle.
