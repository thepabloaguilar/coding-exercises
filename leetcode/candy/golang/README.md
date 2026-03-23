# Candy - Go Solution

## Algorithm Strategy

This solution uses a **two-pass greedy approach** to distribute candies based on rating comparisons.

The algorithm:
1. Initializes a `candies` slice with 1 candy for each child
2. **First pass (left to right - combined with initialization)**:
   - For each child from index 0 to n-1
   - If the current child's rating is higher than the previous child's
   - Gives them one more candy than the previous child
3. **Second pass (right to left)**:
   - For each child from index n-1 down to 0
   - If the current child's rating is higher than the next child's
   - Updates their candies to max(current, next + 1)
   - Simultaneously accumulates the sum in `candiesSum`
4. Returns the total sum of candies

The key insight is using two passes to satisfy constraints in both directions, ensuring higher-rated children have more candies than adjacent lower-rated children.

## Time Complexity

**O(n)** where n is the number of children.

The algorithm makes two passes through the ratings array: one forward and one backward. Each pass is O(n).

## Space Complexity

**O(n)** for the candies array.

An additional array of size n stores the candy distribution.
