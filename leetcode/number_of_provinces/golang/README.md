# Number of Provinces - Go Solution

## Algorithm Strategy

This solution uses **DFS to count connected components** in an adjacency matrix.

The algorithm:
1. Initializes `numberOfProvinces` to 0 and boolean `visited` array
2. For each city index:
   - If not visited:
     - Increments province counter (new component found)
     - Runs DFS to mark all connected cities
3. Returns total provinces

The `dfs` helper:
- Marks current city as visited
- Iterates through adjacency matrix row
- For each connected (value == 1) unvisited city:
   - Recursively visits it

The key insight is graph connectivity: each province is a connected component. Cities in the same province can reach each other through direct or indirect connections.

## Time Complexity

**O(n²)** where n is the number of cities.

Main loop is O(n), and each DFS can visit O(n) cities, checking the adjacency matrix.

## Space Complexity

**O(n)** for visited array and recursion stack.

The visited array has size n. Maximum recursion depth is n for a fully connected component.
