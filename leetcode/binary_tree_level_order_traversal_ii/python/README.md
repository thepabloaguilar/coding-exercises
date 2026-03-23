# Binary Tree Level Order Traversal II - Python Solution

## Algorithm Strategy

This solution uses **Breadth-First Search (BFS)** with a queue, then reverses the result to get bottom-up order.

The algorithm:
1. Returns an empty list if the root is None
2. Initializes a deque with the root node and an empty `values` list
3. While the queue has elements:
   - Creates `level_values` for the current level
   - Processes all nodes at the current level (using `len(queue)`)
   - For each node: pops from left, adds value, enqueues children
   - Appends the level to `values`
4. Reverses the entire result using slicing `values[::-1]` to get bottom-up order
5. Returns the reversed list

The key insight is to perform standard level-order traversal (top to bottom) and then reverse the entire result array to achieve bottom-up ordering.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Each node is visited once during BFS traversal. The reversal operation is also O(n) as it processes the entire result list.

## Space Complexity

**O(w + n)** where w is the maximum width of the tree and n is the total number of nodes.

The queue holds at most one complete level (O(w)). The output list stores all node values (O(n)). In a complete binary tree, w can be approximately n/2.
