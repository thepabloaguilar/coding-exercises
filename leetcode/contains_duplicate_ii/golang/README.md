# Contains Duplicate II - Go Solution

## Algorithm Strategy

This solution uses a **hash map** to track the most recent index of each number encountered.

The algorithm:
1. Initializes an empty map `seenNumbers` to store number-to-index mappings
2. Iterates through the array with index and value
3. For each number:
   - Checks if the number exists in the map
   - If found and the distance (idx - seenIdx) is ≤ k:
     - Returns true immediately (found nearby duplicate)
   - Updates the map with the current index for this number
4. If no valid duplicate pair is found, returns false

The key insight is maintaining only the most recent occurrence of each number. When we encounter a number again, we only need to check if it's within distance k from its last appearance.

## Time Complexity

**O(n)** where n is the length of the nums array.

Single pass through the array with O(1) map lookup and insertion for each element.

## Space Complexity

**O(min(n, k))** in practice, but O(n) in the worst case.

The map stores at most n unique numbers. When duplicates exist within distance k, the effective size is constrained by the window.
