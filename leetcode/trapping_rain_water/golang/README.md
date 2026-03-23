# Trapping Rain Water - Go Solution

## Algorithm Strategy

This solution uses the **two-pointer technique** with running max heights for optimal space.

The algorithm:
1. Initializes left and right pointers at both ends
2. Tracks leftMax and rightMax (tallest bars seen from each side)
3. While left < right:
   - If leftMax <= rightMax:
     - Moves left pointer, updates leftMax
     - Adds water: leftMax - height[left]
   - Otherwise:
     - Moves right pointer, updates rightMax
     - Adds water: rightMax - height[right]
4. Returns total water

The key insight: move from the side with smaller max. The water level is limited by the smaller of the two max heights. We can calculate water at the current position using the smaller max.

## Time Complexity

**O(n)** where n is the array length.

Single pass with two pointers converging from ends.

## Space Complexity

**O(1)** constant space.

Only a few variables for pointers and max values.
