# Majority Element - Python Solution

## Algorithm Strategy

This solution uses a **hash map (defaultdict) with early termination**.

The algorithm:
1. Calculates the threshold as `len(nums) // 2`
2. Initializes a defaultdict to count occurrences
3. For each number in the array:
   - Increments its count in the dictionary
   - If count exceeds threshold, immediately returns the number
4. Raises an exception if no majority element found (though problem guarantees one exists)

The key insight is early termination: as soon as we find an element that appears more than n/2 times, we can return immediately without processing the rest of the array.

## Time Complexity

**O(n)** where n is the length of the nums array.

In the worst case, we iterate through the entire array. Dictionary operations are O(1) average case.

## Space Complexity

**O(n)** for the hash map in the worst case.

The dictionary could store up to n unique elements, though in practice it will be smaller since a majority element exists.
