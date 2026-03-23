# Binary Tree Level Order Traversal - Python Solution

## Algorithm Strategy

This solution uses **Breadth-First Search (BFS)** with a queue to traverse the tree level by level.

The algorithm:
1. Returns an empty list if the root is None
2. Initializes an empty `values` list to store level-by-level results
3. Creates a deque (double-ended queue) with the root node
4. While the queue is not empty:
   - Creates a `level_values` list for the current level
   - Processes all nodes in the current level (using `len(queue)` to know the count)
   - For each node: pops it from the left, adds its value to `level_values`
   - Enqueues left and right children (if they exist) for the next level
   - Appends the completed `level_values` to the result
5. Returns the list of levels

The key insight is using the queue's current length to process exactly one level at a time, ensuring proper level separation.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Each node is visited exactly once, and for each node, we perform constant-time enqueue/dequeue operations.

## Space Complexity

**O(w)** where w is the maximum width of the tree.

The queue stores at most one complete level at a time. In the worst case (a complete binary tree), the last level could contain roughly n/2 nodes. The output list also uses O(n) space to store all node values.
