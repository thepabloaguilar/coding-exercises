# Ransom Note - Python Solution

## Algorithm Strategy

This solution uses a **character frequency counter** with early termination.

The algorithm:
1. Returns False if ransomNote is longer than magazine (impossible)
2. Creates a defaultdict counter
3. Counts all characters in magazine
4. For each character in ransomNote:
   - Decrements its count
   - If count becomes negative, returns False immediately (insufficient characters)
5. Returns True if all characters were available

The key insight is using character frequency: we need at least as many of each character in magazine as required by ransomNote.

## Time Complexity

**O(m + n)** where m is magazine length and n is ransomNote length.

Counting magazine characters is O(m), checking ransomNote is O(n) with early termination possible.

## Space Complexity

**O(k)** where k is the number of unique characters in magazine.

The dictionary stores at most 26 entries for lowercase English letters.
