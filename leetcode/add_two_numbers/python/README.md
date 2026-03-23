# Add Two Numbers - Python Solution

## Algorithm Strategy

This solution uses a **single-pass linked list traversal** approach to add two numbers represented as linked lists in reverse order.

The algorithm:
1. Initializes a dummy head node and a pointer to build the result list
2. Maintains a `rest` variable to track the carry from digit addition
3. Iterates through both linked lists simultaneously using a while loop
4. For each position, extracts values from both lists (or 0 if a list has ended)
5. Calculates the sum including any carry from the previous position
6. Stores the digit (sum % 10) in a new node and updates the carry (sum // 10)
7. Advances pointers in both input lists
8. After the main loop, creates a final node if there's a remaining carry
9. Returns the result starting from dummy head's next node

## Time Complexity

**O(max(m, n))** where m and n are the lengths of the two input linked lists.

The algorithm makes a single pass through both lists, processing each node exactly once. The iteration continues until both lists are fully traversed.

## Space Complexity

**O(max(m, n))** for the output linked list.

The solution creates a new linked list to store the result, which in the worst case (with a carry) can have max(m, n) + 1 nodes. The algorithm uses only a constant amount of additional space O(1) for variables like `rest`, `first_number`, and `second_number`.
