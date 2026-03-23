# Symmetric Tree - Python Solution

## Algorithm Strategy

This solution uses **recursive mirror comparison** of left and right subtrees.

The algorithm:
1. Returns True if root is None
2. Calls _traverse with root's left and right children
3. _traverse logic:
   - Both None: returns True (symmetric nulls)
   - One None: returns False (asymmetric structure)
   - Values differ: returns False
   - Recursively checks:
     - node_a.left vs node_b.right (mirror positions)
     - node_a.right vs node_b.left (mirror positions)
4. Returns True only if all checks pass

The key insight is mirror traversal: for symmetry, left subtree's left should match right subtree's right, and left subtree's right should match right subtree's left.

## Time Complexity

**O(n)** where n is the number of nodes.

Visits each node at most once during recursive traversal.

## Space Complexity

**O(h)** where h is the tree height.

Recursion stack proportional to tree height. O(log n) for balanced, O(n) for skewed.
