# Product of Array Except Self - Go Solution

## Algorithm Strategy

This solution uses **two-pass prefix/postfix product accumulation** without division.

The algorithm:
1. Creates result slice initialized with zeros
2. **First pass (forward - prefix)**:
   - Accumulates prefix product (product of all elements before current index)
   - Stores prefix in result[idx]
   - Updates prefix by multiplying current element
3. **Second pass (backward - postfix)**:
   - Accumulates postfix product (product of all elements after current index)
   - Multiplies result[idx] by postfix
   - Updates postfix by multiplying current element
4. Returns result slice

The key insight: product except self at index i = (product of elements before i) × (product of elements after i). Two passes compute these products separately.

## Time Complexity

**O(n)** where n is array length.

Two sequential O(n) passes through the array.

## Space Complexity

**O(1)** excluding the output array.

Only constant extra space for prefix and postfix variables. The result array is required output.
