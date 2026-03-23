# Maximum Depth of Binary Tree - Python Solution

## Algorithm Strategy

This solution uses **recursive depth-first search (DFS)** to calculate tree depth.

The algorithm:
1. The main method calls the helper `_dfs` method with the root
2. The `_dfs` method:
   - Takes a node and current level (default 0)
   - Returns level if node is None (base case)
   - Recursively calls _dfs on left child with level + 1
   - Recursively calls _dfs on right child with level + 1
   - Returns the maximum of both recursive calls
3. Returns the maximum depth found

The key insight is that the depth of a tree is 1 + the maximum depth of its left and right subtrees. The level parameter tracks the current depth as we recurse.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Every node is visited exactly once during the DFS traversal.

## Space Complexity

**O(h)** where h is the height of the tree.

The recursion stack uses space proportional to tree height. For a skewed tree, h = n (O(n)). For a balanced tree, h = log n (O(log n)).
