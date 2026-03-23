# Trapping Rain Water - Python Solution

## Algorithm Strategy

This solution uses **precomputed max heights** on both sides for each position.

The algorithm:
1. Creates `max_left` array storing maximum height to the left of each index
2. Forward pass: tracks running maximum from left
3. Creates `max_right` array storing maximum height to the right of each index
4. Backward pass: tracks running maximum from right
5. For each position:
   - Water level = min(max_left, max_right) - height[idx]
   - If positive, adds to result
6. Returns total water trapped

The key insight: water at position i is determined by the minimum of the tallest bars on its left and right, minus the bar's own height.

## Time Complexity

**O(n)** where n is the array length.

Three passes: computing max_left (O(n)), max_right (O(n)), and calculating water (O(n)).

## Space Complexity

**O(n)** for the two auxiliary arrays.

Two arrays of size n store the maximum heights from each direction.
