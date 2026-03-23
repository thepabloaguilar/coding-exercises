# Contains Duplicate II - Python Solution

## Algorithm Strategy

This solution uses a **hash map** to track the most recent index of each number.

The algorithm:
1. Initializes an empty dictionary `seen_numbers` to map values to their most recent indices
2. Iterates through the array with both index and value
3. For each number:
   - Checks if the number has been seen before using `get(num)`
   - If seen and the distance between current index and previous index is ≤ k:
     - Returns True (found a nearby duplicate)
   - Otherwise, updates the dictionary with the current index
4. If no valid duplicate pair is found, returns False

The key insight is that we only need to track the most recent occurrence of each number. If we find the same number again within distance k from its last occurrence, we have our answer.

## Time Complexity

**O(n)** where n is the length of the nums array.

The algorithm makes a single pass through the array with constant-time dictionary operations (get and set) for each element.

## Space Complexity

**O(min(n, k))** in practice, but O(n) in the worst case.

The dictionary stores at most n unique numbers. In practice, if duplicates exist within distance k, the dictionary size is limited by the sliding window constraint.
