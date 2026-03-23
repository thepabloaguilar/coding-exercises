# Container With Most Water - Python Solution

## Algorithm Strategy

This solution uses the **two-pointer technique** to find the maximum container area.

The algorithm:
1. Initializes `left` pointer at start (0) and `right` pointer at end (n-1)
2. Calculates initial container area using the minimum of the two heights and the distance between them
3. While left < right:
   - Moves the pointer with the smaller height inward (left++ or right--)
   - This is optimal because the limiting factor is the shorter line
   - Calculates new area and updates `maxContainer` if larger
4. Returns the maximum area found

The key insight is that moving the pointer with the shorter height is optimal because:
- The width will always decrease as pointers move inward
- Moving the taller pointer can never increase area (limited by shorter height)
- Moving the shorter pointer gives a chance to find a taller line

## Time Complexity

**O(n)** where n is the length of the height array.

The algorithm makes a single pass with two pointers converging from opposite ends, visiting each element at most once.

## Space Complexity

**O(1)** constant space.

Only a few variables (left, right, maxContainer) are used regardless of input size.
