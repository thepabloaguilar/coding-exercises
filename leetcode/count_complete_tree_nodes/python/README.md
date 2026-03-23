# Count Complete Tree Nodes - Python Solution

## Algorithm Strategy

This solution uses **optimized recursion** that exploits the complete binary tree property.

The algorithm:
1. Returns 0 if root is None
2. Calculates the left depth (always going left from left child)
3. Calculates the right depth (always going right from right child)
4. If left depth equals right depth:
   - The tree is a perfect binary tree
   - Returns (2^leftDepth) - 1 using the formula for perfect tree node count
5. Otherwise:
   - Recursively counts nodes in left and right subtrees
   - Returns 1 (root) + count(left) + count(right)

The helper methods:
- `_getLeftDepth`: Goes left from each node to measure the leftmost path
- `_getRightDepth`: Goes right from each node to measure the rightmost path

The key insight is that in a complete binary tree, if the leftmost and rightmost paths from a node have the same depth, that subtree is perfect and we can use the formula instead of counting each node.

## Time Complexity

**O(log²n)** where n is the number of nodes.

In the worst case, we traverse O(log n) depth and at each level make O(log n) depth measurements. This gives O(log²n) total operations, which is better than O(n) for a regular tree traversal.

## Space Complexity

**O(log n)** for the recursion stack.

The maximum recursion depth is the height of the tree, which is O(log n) for a complete binary tree.
