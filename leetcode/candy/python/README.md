# Candy - Python Solution

## Algorithm Strategy

This solution uses a **two-pass greedy approach** to ensure each child gets appropriate candies based on ratings.

The algorithm:
1. Initializes a `candies` array with 1 candy for each child
2. **First pass (left to right)**:
   - For each child from index 1 to n-1
   - If the current child's rating is higher than the previous child's
   - Gives them one more candy than the previous child
3. **Second pass (right to left)**:
   - For each child from index n-2 down to 0
   - If the current child's rating is higher than the next child's
   - Updates their candies to be the maximum of current value and (next child's candies + 1)
4. Returns the sum of all candies

The key insight is that two passes ensure both local constraints are satisfied: higher-rated children have more candies than both their left and right neighbors.

## Time Complexity

**O(n)** where n is the number of children.

The algorithm makes two complete passes through the ratings array (one forward, one backward), and then sums the candies array. Each operation is O(n).

## Space Complexity

**O(n)** for the candies array.

An additional array of size n is used to store the candy count for each child.
