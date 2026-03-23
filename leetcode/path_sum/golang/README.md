# Path Sum - Go Solution

## Algorithm Strategy

This solution uses **recursive DFS** with decreasing target sum.

The algorithm:
1. Returns false if root is nil
2. Decrements targetSum by current node's value
3. If node is a leaf (both children nil):
   - Returns true if targetSum equals 0 (valid path found)
4. Otherwise, recursively checks both subtrees with updated targetSum
5. Returns true if either subtree contains a valid path (OR logic)

The key insight is maintaining a running difference: as we traverse down, we subtract each node's value. Upon reaching a leaf, if the remaining sum is exactly 0, we've found a valid root-to-leaf path matching the target.

## Time Complexity

**O(n)** where n is the number of nodes.

Each node is visited at most once during the recursive traversal.

## Space Complexity

**O(h)** where h is the tree height.

Recursion stack depth equals tree height. Worst case O(n) for skewed tree, O(log n) for balanced tree.
