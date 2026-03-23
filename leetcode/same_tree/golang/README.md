# Same Tree - Go Solution

## Algorithm Strategy

This solution uses **recursive parallel tree traversal**.

The algorithm:
1. Base cases:
   - Both nil: returns true (matching)
   - One nil: returns false (structure differs)
2. Value comparison:
   - If p.Val != q.Val: returns false
3. Recursively checks both subtrees:
   - Left subtrees must match AND
   - Right subtrees must match
4. Returns true only if all conditions satisfied

The key insight is simultaneous recursive traversal: at each step, compare current nodes and recursively verify their subtrees match.

## Time Complexity

**O(min(n, m))** where n and m are node counts.

Visits all nodes if trees are identical, stops early if mismatch found.

## Space Complexity

**O(min(h1, h2))** for recursion stack.

Stack depth limited by the shorter tree's height.
