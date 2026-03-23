# Insert Delete GetRandom O(1) - Python Solution

## Algorithm Strategy

This solution uses a **combination of dictionary and list** to achieve O(1) operations.

The data structures:
- `_dict`: Maps values to their indices in the list
- `_list`: Stores all values in insertion order

The operations:
1. **Insert**:
   - Checks if value exists in dictionary
   - If not, appends to list and stores index in dictionary
   - Returns whether insertion occurred
2. **Remove**:
   - Checks if value exists
   - If yes, swaps the target value with the last element in the list
   - Updates the swapped element's index in the dictionary
   - Removes the last element and deletes from dictionary
   - Returns whether removal occurred
3. **GetRandom**:
   - Uses `random.choice()` on the list for uniform random selection

The key insight is using the swap-with-last technique for O(1) removal while maintaining a contiguous array for random access.

## Time Complexity

**O(1)** average time for all operations (insert, remove, getRandom).

Dictionary operations (lookup, insert, delete) are O(1) average. List append and pop from end are O(1). Random choice from list is O(1).

## Space Complexity

**O(n)** where n is the number of elements.

Both the dictionary and list store all n elements.
