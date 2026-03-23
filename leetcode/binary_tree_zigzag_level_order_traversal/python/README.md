# Binary Tree Zigzag Level Order Traversal - Python Solution

## Algorithm Strategy

This solution uses **Breadth-First Search (BFS)** with alternating direction control for zigzag ordering.

The algorithm:
1. Returns an empty list if root is None
2. Initializes a deque with root, empty `values` list, and `to_left` flag set to True
3. While the queue has elements:
   - Creates `level_values` for the current level
   - Processes all nodes at the current level
   - For each node: pops from left
     - If `to_left` is True: appends value to the end
     - If `to_left` is False: inserts value at position 0 (reverses order)
   - Enqueues left and right children
   - Toggles `to_left` flag for the next level
   - Appends completed level to results
4. Returns the zigzag-ordered result

The key insight is using `insert(0, val)` to build the level in reverse order on alternate levels, achieving zigzag ordering without explicit reversal.

## Time Complexity

**O(n^2)** in the worst case, where n is the number of nodes.

While each node is visited once (O(n)), the `insert(0, val)` operation is O(k) where k is the level size. In a complete binary tree with width w, this leads to O(w^2) per level. Summing across all levels can approach O(n^2) in the worst case.

## Space Complexity

**O(w)** where w is the maximum width of the tree.

The queue holds at most one complete level at a time. The output list requires O(n) space for all node values.
