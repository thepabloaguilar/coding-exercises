# Binary Tree Level Order Traversal II - Go Solution

## Algorithm Strategy

This solution uses **Breadth-First Search (BFS)** with a queue, then reverses the result using `slices.Reverse`.

The algorithm:
1. Returns an empty slice if root is nil
2. Initializes a queue with the root node and an empty `values` slice
3. While the queue has elements:
   - Captures the current level size
   - Pre-allocates `levelValues` with appropriate capacity
   - Processes all nodes at current level (indices 0 to size-1)
   - For each node: adds value, enqueues children
   - Appends level to results and removes processed nodes
4. Reverses the entire result using `slices.Reverse(values)` for bottom-up order
5. Returns the reversed slice

The key insight is performing standard top-down level-order traversal and then reversing the complete result to achieve bottom-up ordering.

## Time Complexity

**O(n)** where n is the number of nodes in the tree.

Each node is visited exactly once during BFS. The `slices.Reverse` operation is O(k) where k is the number of levels, which is at most O(log n) for balanced trees or O(n) for skewed trees.

## Space Complexity

**O(w + n)** where w is the maximum width and n is the total number of nodes.

The queue stores at most one complete level (O(w)). The output slice stores all node values (O(n)).
