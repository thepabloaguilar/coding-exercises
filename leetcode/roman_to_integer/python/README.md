# Roman to Integer - Python Solution

## Algorithm Strategy

This solution uses a **character-by-character accumulation** with subtraction correction.

The algorithm:
1. Creates values_mapping dictionary for Roman numerals
2. Iterates through each character:
   - Gets its integer value
   - If last_value exists and current > last_value (subtractive notation like IV):
     - Corrects by subtracting twice the last_value and adding difference
     - Resets last_value to None
   - Otherwise: adds value and updates last_value
3. Returns accumulated result

The key insight is handling subtractive notation (IV, IX, XL, XC, CD, CM): when a smaller value precedes a larger one, the smaller is subtracted. The correction formula accounts for having already added the smaller value.

## Time Complexity

**O(n)** where n is the string length.

Single pass through the string with constant-time dictionary lookups.

## Space Complexity

**O(1)** constant space.

The dictionary has fixed size (7 entries) regardless of input.
