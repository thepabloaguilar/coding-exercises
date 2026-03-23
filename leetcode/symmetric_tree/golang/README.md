# Symmetric Tree - Go Solution

## Algorithm Strategy

This solution uses **recursive mirror traversal** to verify symmetry.

The algorithm:
1. Returns true if root is nil
2. Calls traverse helper with left and right children
3. traverse function:
   - Both nil: returns true (symmetric)
   - One nil: returns false (asymmetric)
   - Values differ: returns false
   - Recursively verifies:
     - nodeA.Left mirrors nodeB.Right
     - nodeA.Right mirrors nodeB.Left
4. Returns true if all mirror pairs match

The key insight is checking mirror positions: for a symmetric tree, the left subtree is a mirror reflection of the right subtree. This requires comparing crossed positions (left-right and right-left).

## Time Complexity

**O(n)** where n is the number of nodes.

Each node visited once during traversal.

## Space Complexity

**O(h)** where h is the tree height.

Recursion stack depth equals tree height.
