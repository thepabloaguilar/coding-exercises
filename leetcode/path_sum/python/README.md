# Path Sum - Python Solution

## Algorithm Strategy

This solution uses **recursive DFS** with running sum subtraction.

The algorithm:
1. Returns False if root is None
2. Subtracts current node's value from targetSum
3. If at a leaf node (no children):
   - Returns True if targetSum equals 0 (path found)
4. Otherwise, recursively checks left and right subtrees with updated targetSum
5. Returns True if either subtree has a valid path (OR operation)

The key insight is working with a decreasing target: subtract each node's value as we traverse. When we reach a leaf, check if the remaining target is 0, meaning we found a valid root-to-leaf path with the exact sum.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

In the worst case (no valid path), we visit every node once.

## Space Complexity

**O(h)** where h is the height of the tree.

The recursion stack uses space proportional to tree height. Worst case O(n) for skewed tree, O(log n) for balanced tree.
