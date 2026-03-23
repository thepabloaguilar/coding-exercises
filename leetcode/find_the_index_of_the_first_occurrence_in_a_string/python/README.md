# Find the Index of the First Occurrence in a String - Python Solution

## Algorithm Strategy

This solution uses Python's **built-in `find()` method** for string searching.

The algorithm:
1. Directly calls `haystack.find(needle)`
2. The `find()` method returns:
   - The index of the first occurrence of `needle` in `haystack`, or
   - -1 if `needle` is not found

Python's `find()` method typically uses an optimized string searching algorithm (likely Boyer-Moore or similar) implemented in C.

## Time Complexity

**O(n × m)** in the worst case, where n is the length of haystack and m is the length of needle.

The built-in `find()` method has worst-case O(n × m) for simple implementations, though Python's actual implementation may use optimizations. In practice, it performs much better on average.

## Space Complexity

**O(1)** constant space.

The solution uses no additional data structures, only the built-in string method.
