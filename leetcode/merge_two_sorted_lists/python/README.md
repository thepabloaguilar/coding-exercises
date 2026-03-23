# Merge Two Sorted Lists - Python Solution

## Algorithm Strategy

This solution uses a **two-pointer merge** approach with a dummy head node.

The algorithm:
1. Creates a dummy head node and tracks current position
2. While either list has nodes:
   - Extracts values from both lists (uses 101 as sentinel for exhausted lists)
   - Compares values and creates new node with the smaller value
   - Advances the pointer of the list that contributed the value
   - Moves the result pointer forward
3. Returns head.next (skipping dummy node)

The key insight is using a sentinel value (101, which exceeds the problem's constraint) to simplify the comparison logic when one list is exhausted.

## Time Complexity

**O(m + n)** where m and n are the lengths of the two lists.

Every node from both lists is visited exactly once.

## Space Complexity

**O(m + n)** for the new linked list.

A new node is created for each element in the merged result.
