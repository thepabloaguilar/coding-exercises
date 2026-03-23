# Rotate Array - Go Solution

## Algorithm Strategy

This solution uses the **triple reverse algorithm** for in-place rotation.

The algorithm:
1. Normalizes k using modulo
2. Performs three reversals:
   - Reverse entire array
   - Reverse first k elements
   - Reverse remaining n-k elements
3. Array is now rotated right by k positions

Example for nums=[1,2,3,4,5,6,7], k=3:
- After reversing all: [7,6,5,4,3,2,1]
- After reversing first 3: [5,6,7,4,3,2,1]
- After reversing last 4: [5,6,7,1,2,3,4]

The key insight is that three reversals achieve rotation without extra space.

## Time Complexity

**O(n)** where n is the array length.

Each of the three reversals is O(n).

## Space Complexity

**O(1)** constant space.

In-place reversals using slices.Reverse with no additional arrays.
