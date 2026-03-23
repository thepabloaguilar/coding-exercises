# Container With Most Water - Go Solution

## Algorithm Strategy

This solution uses the **two-pointer technique** to efficiently find the maximum water container area.

The algorithm:
1. Initializes `left` pointer at index 0 and `right` pointer at the last index
2. Calculates initial area: min(height[left], height[right]) × (right - left)
3. While left < right:
   - Moves the pointer pointing to the shorter line inward
   - If height[left] < height[right]: increment left
   - Otherwise: decrement right
   - Calculates new area and updates `maxContainer` if it's larger
4. Returns the maximum area found

The key insight is the greedy approach: always move the pointer with the shorter height because:
- Width decreases with each step
- The shorter height limits the container's capacity
- Moving the taller pointer cannot increase area
- Only by moving the shorter pointer might we find a taller line

## Time Complexity

**O(n)** where n is the length of the height array.

Single pass with two pointers converging from both ends, each element examined at most once.

## Space Complexity

**O(1)** constant space.

Only three variables (left, right, maxContainer) are used.
