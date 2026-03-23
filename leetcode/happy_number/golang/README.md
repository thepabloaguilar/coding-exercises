# Happy Number - Go Solution

## Algorithm Strategy

This solution uses **cycle detection with a hash map** to identify happy numbers.

The algorithm:
1. Initializes `seenNumbers` map with the initial value n
2. While n is not 1:
   - Calculates sum of squares of digits using `sumOfSquares` helper
   - Checks if this sum exists in the map:
     - If yes, we've found a cycle, returns false
   - Adds the sum to seenNumbers
   - Updates n to this sum
3. If we reach 1, returns true

The helper function `sumOfSquares`:
- Extracts each digit using modulo 10
- Squares it using `math.Pow` and adds to result
- Divides by 10 to process the next digit
- Continues until n becomes 0

The key insight is that non-happy numbers eventually cycle. By tracking all visited numbers in a map, we can detect when we enter a cycle.

## Time Complexity

**O(log n)** where n is the input number.

The number of digits is O(log n). Each iteration typically reduces the number's magnitude. The map ensures we detect cycles in constant time.

## Space Complexity

**O(log n)** for the map storing seen numbers.

The map size is bounded by the number of unique values before reaching 1 or a cycle.
