# Linked List Cycle - Python Solution

## Algorithm Strategy

This solution uses a **hash set** to track visited nodes and detect cycles.

The algorithm:
1. Initializes an empty set `seen` to store visited nodes
2. Starts with the head node
3. While the current node exists:
   - Checks if the node is already in the `seen` set
   - If yes, returns True (cycle detected)
   - Otherwise, adds the node to `seen` and moves to next node
4. If we reach the end (None), returns False (no cycle)

The key insight is that in a cycle, we'll eventually revisit a node. By storing node references (not values) in a set, we can detect when we encounter the same node object again.

## Time Complexity

**O(n)** where n is the number of nodes in the list.

Each node is visited at most once. Set operations (add and lookup) are O(1) average case.

## Space Complexity

**O(n)** for the hash set.

In the worst case (no cycle), we store all n nodes in the set.
