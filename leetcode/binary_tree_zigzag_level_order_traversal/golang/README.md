# Binary Tree Zigzag Level Order Traversal - Go Solution

## Algorithm Strategy

This solution uses **Breadth-First Search (BFS)** with alternating prepend/append logic for zigzag ordering.

The algorithm:
1. Returns an empty slice if root is nil
2. Initializes queue with root, empty `values` slice, and `toLeft` flag set to true
3. While the queue has elements:
   - Captures current level size
   - Pre-allocates `levelValues` slice
   - Processes all nodes at current level
   - For each node:
     - If `toLeft` is true: appends value normally
     - If `toLeft` is false: prepends value using `append([]int{node.Val}, levelValues...)`
   - Enqueues left and right children
   - Toggles `toLeft` for next level
   - Appends completed level to results and removes processed nodes
4. Returns the zigzag-ordered result

The key insight is using slice prepend operations on alternate levels to achieve zigzag ordering without explicit reversal.

## Time Complexity

**O(n^2)** in the worst case, where n is the number of nodes.

Each node is visited once, but the prepend operation `append([]int{val}, slice...)` creates a new slice and copies all elements, which is O(k) for a level of size k. This can sum to O(n^2) across all levels in a complete tree.

## Space Complexity

**O(w)** where w is the maximum width of the tree.

The queue stores at most one complete level. The output slice requires O(n) space for all node values.
