# Merge Two Sorted Lists - Go Solution

## Algorithm Strategy

This solution uses a **two-pointer merge** with a dummy head node.

The algorithm:
1. Creates dummy head node and current pointer
2. While either list has nodes remaining:
   - Sets num1 to list1's value or sentinel 101 if list1 is nil
   - Sets num2 to list2's value or sentinel 101 if list2 is nil
   - Creates new node with smaller value
   - Advances corresponding list pointer
   - Moves current pointer forward
3. Returns head.Next (skipping dummy)

The key insight is using sentinel value 101 (exceeds problem constraints) to handle exhausted lists uniformly, simplifying comparison logic.

## Time Complexity

**O(m + n)** where m and n are the list lengths.

Each node from both lists is processed exactly once.

## Space Complexity

**O(m + n)** for the new linked list.

Creates a new node for each element in the merged result.
