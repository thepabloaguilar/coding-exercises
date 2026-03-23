# Find the Index of the First Occurrence in a String - Go Solution

## Algorithm Strategy

This solution uses a **brute-force substring matching** approach with nested loops.

The algorithm:
1. Returns 0 if needle is empty (edge case)
2. Returns -1 if needle is longer than haystack (impossible match)
3. Iterates through haystack from index 0 to len(haystack) - len(needle)
4. For each starting position in haystack:
   - Compares characters of haystack and needle one by one
   - If a mismatch is found, breaks and tries the next starting position
   - If all characters match (reaches last character of needle), returns the starting index
5. Returns -1 if no match is found

The key insight is checking each possible starting position in haystack and comparing character by character with needle.

## Time Complexity

**O(n × m)** where n is the length of haystack and m is the length of needle.

In the worst case (many near-matches), we may compare up to m characters for each of the n-m+1 starting positions.

## Space Complexity

**O(1)** constant space.

Only a few variables are used for indexing, regardless of input size.
