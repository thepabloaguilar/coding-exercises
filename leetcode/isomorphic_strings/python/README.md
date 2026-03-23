# Isomorphic Strings - Python Solution

## Algorithm Strategy

This solution uses a **dictionary for character mapping and a set to track used characters**.

The algorithm:
1. Returns False if strings have different lengths
2. Initializes `mapping` dictionary and `seen_letters` set
3. Iterates through both strings simultaneously by index:
   - Gets the character from s and checks if it has a mapping
   - If no mapping exists:
     - Checks if t[idx] is already mapped to another character
     - If yes, returns False (two characters can't map to the same character)
     - Otherwise, creates the mapping and marks t[idx] as used
   - If mapping exists but doesn't match t[idx], returns False
4. If all characters map correctly, returns True

The key insight is that we need to ensure:
1. Each character in s maps to exactly one character in t
2. No two characters in s map to the same character in t (bijection)

## Time Complexity

**O(n)** where n is the length of the strings.

Single pass through the strings with constant-time dictionary and set operations.

## Space Complexity

**O(k)** where k is the number of unique characters (at most 128 for ASCII or 256 for extended ASCII).

The dictionary and set store mappings for unique characters.
