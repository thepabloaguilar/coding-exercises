# Linked List Cycle - Go Solution

## Algorithm Strategy

This solution uses a **hash map** to track visited nodes and detect cycles.

The algorithm:
1. Initializes a map `seen` to store visited node pointers
2. Starts with the head node
3. While the current node is not nil:
   - Checks if the node pointer exists in the map
   - If yes, returns true (cycle detected)
   - Otherwise, marks the node as seen and advances to next node
4. If we reach nil, returns false (no cycle exists)

The key insight is storing node pointers (memory addresses) in the map. If we encounter the same pointer twice, we've found a cycle.

## Time Complexity

**O(n)** where n is the number of nodes in the list.

Each node is visited once, with O(1) average-case map operations.

## Space Complexity

**O(n)** for the hash map.

In the worst case (no cycle), all n node pointers are stored in the map.
