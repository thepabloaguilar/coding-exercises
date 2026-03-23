# Number of Provinces - Python Solution

## Algorithm Strategy

This solution uses **DFS to find connected components** in an adjacency matrix graph.

The algorithm:
1. Initializes `number_of_components` to 0 and `visited` set
2. For each city (node) from 0 to size-1:
   - If not yet visited:
     - Increments component counter (found new province)
     - Runs DFS from this city to mark all connected cities
3. Returns total number of components

The `_dfs` helper:
- Marks current city as visited
- Iterates through all cities
- If connected (is_connected[node_idx][idx] == 1) and not visited:
   - Recursively visits that city

The key insight is treating this as a graph connectivity problem. Each province is a connected component in the graph where edges represent direct connections between cities.

## Time Complexity

**O(n²)** where n is the number of cities.

The outer loop runs n times, and DFS potentially checks all n cities for each unvisited city. The adjacency matrix has n² entries.

## Space Complexity

**O(n)** for the visited set and recursion stack.

The visited set stores up to n cities. Maximum recursion depth is n in a fully connected graph.
