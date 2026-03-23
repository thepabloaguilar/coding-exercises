# Count Complete Tree Nodes - Go Solution

## Algorithm Strategy

This solution uses **optimized recursion** leveraging the complete binary tree property.

The algorithm:
1. Returns 0 if root is nil
2. Calculates `leftDepth` by traversing leftmost path from left child
3. Calculates `rightDepth` by traversing rightmost path from right child
4. If leftDepth equals rightDepth:
   - The tree is a perfect binary tree
   - Returns (2^leftDepth) - 1 using the perfect tree formula
5. Otherwise:
   - Recursively counts left and right subtrees
   - Returns 1 + CountNodes(left) + CountNodes(right)

The helper functions:
- `getLeftDepth`: Measures depth by going left at each step
- `getRightDepth`: Measures depth by going right at each step

The key insight is that in a complete binary tree, when the leftmost and rightmost depths from a subtree are equal, that entire subtree is perfect. We can calculate its node count using the formula 2^h - 1 instead of visiting every node.

## Time Complexity

**O(log²n)** where n is the number of nodes.

Each recursive call involves O(log n) depth measurements, and there are at most O(log n) recursive calls. This yields O(log²n), significantly better than the O(n) naive approach.

## Space Complexity

**O(log n)** for the recursion call stack.

The maximum recursion depth equals the tree height, which is O(log n) for a complete binary tree.
