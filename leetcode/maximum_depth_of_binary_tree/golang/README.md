# Maximum Depth of Binary Tree - Go Solution

## Algorithm Strategy

This solution uses **recursive depth-first search (DFS)** with level tracking.

The algorithm:
1. The main function calls the helper `dfs` function with root and level 0
2. The `dfs` function:
   - Takes a node and current level
   - Returns level if node is nil (base case - reached beyond leaf)
   - Recursively calls dfs on left child with level + 1
   - Recursively calls dfs on right child with level + 1
   - Returns the maximum of both recursive results
3. Returns the maximum depth found

The key insight is tracking depth via the level parameter: each recursive call increments the level, and we return the maximum level reached when hitting nil nodes.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Each node is visited once during the recursive traversal.

## Space Complexity

**O(h)** where h is the height of the tree.

The call stack depth equals the tree height. Worst case O(n) for skewed tree, O(log n) for balanced tree.
