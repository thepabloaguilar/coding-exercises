# Binary Tree Right Side View - Go Solution

## Algorithm Strategy

This solution uses **Breadth-First Search (BFS)** level-order traversal, extracting the rightmost node from each level.

The algorithm:
1. Returns an empty slice if root is nil
2. Initializes a queue with the root and an empty `values` slice
3. While the queue has elements:
   - Captures the current level size
   - Immediately appends the last node's value `queue[size-1].Val` to results (rightmost node)
   - Processes all nodes at the current level (indices 0 to size-1)
   - For each node: enqueues left and right children
   - Removes processed nodes from queue
4. Returns the list of rightmost values

The key insight is that the right side view shows the last node at each level when viewed from the right. By accessing the last element of the queue before processing, we get the rightmost visible node.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Every node is visited exactly once during the BFS traversal.

## Space Complexity

**O(w)** where w is the maximum width of the tree.

The queue stores at most one complete level. The output slice is O(h) where h is the height, storing one value per level.
