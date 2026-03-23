# Length of Last Word - Python Solution

## Algorithm Strategy

This solution uses **Python's built-in string methods** for a concise approach.

The algorithm:
1. Calls `strip()` to remove leading and trailing whitespace
2. Calls `split(' ')` to split the string into words by spaces
3. Accesses the last element using `[-1]`
4. Returns the length of this last word

This leverages Python's powerful string manipulation methods to solve the problem in one line.

## Time Complexity

**O(n)** where n is the length of the string.

- `strip()` scans the string: O(n)
- `split()` scans the string and creates a list: O(n)
- Accessing last element and getting length: O(1)

## Space Complexity

**O(n)** for the list created by split.

The `split()` method creates a list of all words, which in the worst case (single-character words) could have O(n) elements.
