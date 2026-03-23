# Product of Array Except Self - Python Solution

## Algorithm Strategy

This solution uses a **two-pass prefix and postfix product approach** without division.

The algorithm:
1. Initializes result array with all 1s
2. **First pass (left to right - prefix products)**:
   - Maintains running `prefix` product (starts at 1)
   - For each index, stores the prefix product (product of all elements to the left)
   - Updates prefix by multiplying with current element
3. **Second pass (right to left - postfix products)**:
   - Maintains running `postfix` product (starts at 1)
   - Multiplies each result[idx] by postfix (product of all elements to the right)
   - Updates postfix by multiplying with current element
4. Returns the result array

The key insight is that the product except self = (product of all left elements) × (product of all right elements). Two passes compute these separately and combine them.

## Time Complexity

**O(n)** where n is the length of the array.

Two sequential passes through the array, each O(n).

## Space Complexity

**O(1)** excluding the output array.

The result array is required for output. Only constant extra space (prefix, postfix variables) is used.
