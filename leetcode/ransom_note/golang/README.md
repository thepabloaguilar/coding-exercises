# Ransom Note - Go Solution

## Algorithm Strategy

This solution uses a **character frequency map** with early failure detection.

The algorithm:
1. Returns false if ransomNote longer than magazine
2. Creates map to count character frequencies
3. Counts all characters (runes) in magazine
4. For each character in ransomNote:
   - Decrements its count in map
   - If count goes negative, returns false (not enough of that character)
5. Returns true if all checks pass

The key insight is frequency matching: ransomNote can be constructed if magazine contains at least the required frequency of each character.

## Time Complexity

**O(m + n)** where m and n are the string lengths.

Counting magazine is O(m), verifying ransomNote is O(n) with possible early exit.

## Space Complexity

**O(k)** where k is the number of unique characters.

Map stores at most 26 entries for lowercase letters or more for Unicode.
