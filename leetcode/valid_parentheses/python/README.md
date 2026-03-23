# Valid Parentheses - Python Solution

## Algorithm Strategy

This solution uses a **stack** to match opening and closing brackets.

The algorithm:
1. Creates pairs mapping (closing -> opening brackets)
2. Creates closing set for quick lookup
3. Iterates through each character:
   - If closing bracket:
     - Checks if stack empty or top doesn't match: returns False
     - Pops matching opening bracket
   - If opening bracket: pushes to stack
4. Returns True if stack is empty (all matched)

The key insight is LIFO matching: the most recent unmatched opening bracket should match the current closing bracket. A stack naturally implements this.

## Time Complexity

**O(n)** where n is the string length.

Single pass with O(1) stack operations (append/pop) for each character.

## Space Complexity

**O(n)** for the stack in the worst case.

If string is all opening brackets, stack grows to size n.
