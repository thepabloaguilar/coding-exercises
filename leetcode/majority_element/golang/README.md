# Majority Element - Go Solution

## Algorithm Strategy

This solution uses the **Boyer-Moore Voting Algorithm** for optimal space efficiency.

The algorithm:
1. Initializes `candidate` to -1 and `count` to 0
2. For each number in the array:
   - If count is 0, sets current number as new candidate with count 1
   - If number equals candidate, increments count
   - If number differs from candidate, decrements count
3. Returns the final candidate

The key insight is that the majority element (appearing > n/2 times) will always "survive" the voting process. When count reaches 0, we switch candidates. The majority element's excess occurrences ensure it becomes the final candidate.

## Time Complexity

**O(n)** where n is the length of the nums array.

Single pass through the array with constant-time operations.

## Space Complexity

**O(1)** constant space.

Only two variables (candidate and count) are used regardless of input size.
