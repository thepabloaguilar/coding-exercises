# Binary Tree Right Side View - Python Solution

## Algorithm Strategy

This solution uses **Breadth-First Search (BFS)** level-order traversal, capturing the rightmost node of each level.

The algorithm:
1. Returns an empty list if root is None
2. Initializes a deque with the root node and an empty `values` list
3. While the queue has elements:
   - Before processing the level, captures the rightmost node's value using `queue[len(queue)-1].val`
   - Appends this value to the result (this is the visible node from the right side)
   - Processes all nodes in the current level
   - For each node: pops from left, enqueues left and right children
4. Returns the list of rightmost values

The key insight is that the right side view consists of the last (rightmost) node at each level. By peeking at the last element of the queue before processing the level, we capture the rightmost node.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Each node is visited exactly once during the BFS traversal.

## Space Complexity

**O(w)** where w is the maximum width of the tree.

The queue holds at most one complete level at a time. In a complete binary tree, the maximum width is approximately n/2. The output list is O(h) where h is the height, as we store one value per level.
