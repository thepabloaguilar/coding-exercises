# Rotate Array - Python Solution

## Algorithm Strategy

This solution uses **array slicing and concatenation** for rotation.

The algorithm:
1. Normalizes k using modulo to handle k >= len(nums)
2. Slices array into two parts:
   - Last k elements: nums[len(nums)-k:]
   - First n-k elements: nums[:len(nums)-k]
3. Concatenates them in rotated order
4. Assigns back to nums[:] to modify in-place

The key insight is that rotating right by k is equivalent to moving the last k elements to the front. Python's slice syntax makes this elegant.

## Time Complexity

**O(n)** where n is the array length.

Slicing and concatenation both create new lists, requiring O(n) operations.

## Space Complexity

**O(n)** for the intermediate lists.

Although nums is modified in-place, slicing creates temporary lists.
