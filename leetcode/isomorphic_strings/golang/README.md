# Isomorphic Strings - Go Solution

## Algorithm Strategy

This solution uses **two bidirectional maps** to ensure one-to-one character mapping.

The algorithm:
1. Returns false if strings have different lengths
2. Converts both strings to rune slices (for Unicode support)
3. Initializes two maps: `mappingS` (s→t) and `mappingT` (t→s)
4. Iterates through both strings by index:
   - Checks if character from s has a mapping in mappingS
   - If not mapped:
     - Checks if t[idx] is already mapped in mappingT
     - If yes, returns false (one-to-one constraint violated)
     - Creates bidirectional mappings in both maps
   - If mapped but doesn't match t[idx], returns false
5. If all characters have consistent bidirectional mappings, returns true

The key insight is using two maps to enforce bijection: each character in s must map to exactly one character in t, and vice versa.

## Time Complexity

**O(n)** where n is the length of the strings.

Single iteration with O(1) map operations at each step.

## Space Complexity

**O(k)** where k is the number of unique characters.

Two maps store bidirectional mappings for unique characters.
