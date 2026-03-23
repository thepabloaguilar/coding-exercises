# Valid Parentheses - Go Solution

## Algorithm Strategy

This solution uses a **stack (slice)** for bracket matching.

The algorithm:
1. Creates pairs map (closing -> opening brackets)
2. Initializes empty stack slice
3. For each rune:
   - If closing bracket:
     - Checks if stack empty or top mismatches: returns false
     - Pops from stack (slices off last element)
   - If opening bracket: pushes to stack
4. Returns true if stack is empty

The key insight is stack-based matching: each closing bracket should match the most recent unmatched opening bracket (LIFO property).

## Time Complexity

**O(n)** where n is the string length.

Single pass with O(1) amortized stack operations.

## Space Complexity

**O(n)** for the stack in the worst case.

Maximum stack size is n when all characters are opening brackets.
