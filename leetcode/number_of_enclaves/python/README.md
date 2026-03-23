# Number of Enclaves - Python Solution

## Algorithm Strategy

This solution uses **DFS from boundaries** to mark all land cells connected to the edge.

The algorithm:
1. Calculates total cells and initializes `not_connected` counter
2. Iterates through entire grid:
   - Decrements counter for water cells (0)
   - For boundary land cells (1), runs DFS using `_check`
   - DFS marks connected cells as visited (2) and counts them
   - Subtracts boundary-connected land from counter
3. Returns remaining count (land cells not connected to boundary)

The `_check` helper:
- Returns 0 if out of bounds or not land (!=1)
- Marks current cell as 2 (visited)
- Recursively checks all 4 directions
- Returns count of connected cells

The key insight is that enclaves are land cells NOT reachable from boundaries. By marking all boundary-connected land, the remainder are enclaves.

## Time Complexity

**O(m × n)** where m and n are the grid dimensions.

Each cell is visited at most twice: once in the main loop and potentially once during DFS.

## Space Complexity

**O(m × n)** in the worst case for the recursion stack.

If the entire grid is land connected from boundary, the DFS stack depth could be m × n.
