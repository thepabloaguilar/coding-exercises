# Longest Consecutive Sequence - Python Solution

## Algorithm Strategy

This solution uses a **hash set with intelligent sequence start detection**.

The algorithm:
1. Converts the nums array to a set `seenNums` for O(1) lookup
2. Initializes result to 0
3. For each number in the set:
   - Checks if `num - 1` exists in the set
   - If not, this number is the start of a sequence
   - Counts consecutive numbers by incrementing length while `num + length` exists
   - Updates result with the maximum length found
4. Returns the maximum sequence length

The key insight is only starting to count from numbers that begin a sequence (where num-1 doesn't exist). This prevents redundant counting and ensures O(n) time.

## Time Complexity

**O(n)** where n is the length of the nums array.

Creating the set is O(n). Although there are nested loops, each number is visited at most twice: once in the outer loop and potentially once when counted as part of a sequence. This amortizes to O(n).

## Space Complexity

**O(n)** for the hash set.

The set stores all unique numbers from the input array.
