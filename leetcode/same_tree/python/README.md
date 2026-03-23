# Same Tree - Python Solution

## Algorithm Strategy

This solution uses **recursive simultaneous DFS** on both trees.

The algorithm:
1. Main method calls _dfs helper with both roots
2. _dfs logic:
   - If both nodes are None: returns True (matching nulls)
   - If only one is None: returns False (structure mismatch)
   - If values differ: returns False
   - Recursively checks left subtrees AND right subtrees
3. Returns True only if all checks pass

The key insight is parallel traversal: compare nodes at corresponding positions in both trees simultaneously. Any mismatch in structure or values returns False.

## Time Complexity

**O(min(n, m))** where n and m are the node counts.

In the worst case (identical trees), we visit all nodes. Otherwise, we stop at the first mismatch.

## Space Complexity

**O(min(h1, h2))** where h1 and h2 are the tree heights.

Recursion stack depth is limited by the shorter tree's height.
