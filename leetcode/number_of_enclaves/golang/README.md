# Number of Enclaves - Go Solution

## Algorithm Strategy

This solution uses **DFS from boundary cells** to identify edge-connected land.

The algorithm:
1. Calculates total cells in `notConnected`
2. Iterates through entire grid:
   - Decrements counter for water cells
   - For boundary land cells (edge cells with value 1):
     - Runs DFS to mark all connected cells
     - Subtracts count from notConnected
3. Returns remaining count (enclaved land)

The `dfs` helper:
- Returns 0 if out of bounds or not unvisited land
- Marks cell as 2 (visited)
- Recursively explores 4 directions
- Returns total connected cells

The key insight: enclaves are land cells unreachable from grid boundaries. By DFS from all boundary land cells and marking them as visited, the unmarked land cells are enclaves.

## Time Complexity

**O(m × n)** where m and n are grid dimensions.

Each cell is visited at most twice: main iteration and potentially during DFS.

## Space Complexity

**O(m × n)** worst case for recursion stack.

Maximum DFS depth occurs when entire grid is boundary-connected land.
