# Invert Binary Tree - Go Solution

## Algorithm Strategy

This solution uses **recursive depth-first traversal** to invert the binary tree.

The algorithm:
1. Returns nil if root is nil (base case)
2. Swaps the left and right children using Go's multiple assignment
3. Recursively inverts the left subtree
4. Recursively inverts the right subtree
5. Returns the root node

The key insight is that tree inversion is a recursive operation where each node's children are swapped and then the inversion is applied to both subtrees.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Each node is visited once to perform the swap operation.

## Space Complexity

**O(h)** where h is the height of the tree.

The call stack depth is proportional to the tree height. For a skewed tree, this is O(n). For a balanced tree, it's O(log n).
