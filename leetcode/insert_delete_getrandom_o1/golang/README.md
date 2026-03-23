# Insert Delete GetRandom O(1) - Go Solution

## Algorithm Strategy

This solution uses a **combination of map and slice** to achieve O(1) time complexity.

The data structures:
- `valuesMap`: Maps values to their indices in the slice
- `valuesSlice`: Stores all values

The operations:
1. **Insert**:
   - Checks if value exists in map
   - If not, appends to slice and stores index in map
   - Returns whether insertion succeeded
2. **Remove**:
   - Checks if value exists in map
   - If yes, swaps with the last element in the slice
   - Updates the last element's index in the map
   - Removes last element from slice and deletes from map
   - Returns whether removal succeeded
3. **GetRandom**:
   - Uses random index generation on the slice for uniform distribution

The key insight is the swap-with-last-element technique: to remove an element in O(1), swap it with the last element, then remove the last element (which is O(1) for slices).

## Time Complexity

**O(1)** average time for all operations.

Map operations (lookup, insert, delete) are O(1) average. Slice append and truncation from end are O(1). Random number generation is O(1).

## Space Complexity

**O(n)** where n is the number of elements.

Both the map and slice store all n elements.
