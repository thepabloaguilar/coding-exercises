# Invert Binary Tree - Python Solution

## Algorithm Strategy

This solution uses **recursive depth-first traversal** to invert the binary tree.

The algorithm:
1. Returns None if root is None (base case)
2. Swaps the left and right children of the current node using Python's tuple unpacking
3. Recursively inverts the left subtree
4. Recursively inverts the right subtree
5. Returns the root node

The key insight is that inverting a tree means swapping left and right children at every node, which can be elegantly done recursively.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Every node is visited exactly once to perform the swap operation.

## Space Complexity

**O(h)** where h is the height of the tree.

The recursion stack uses space proportional to the tree's height. In the worst case (skewed tree), this is O(n). For a balanced tree, it's O(log n).
