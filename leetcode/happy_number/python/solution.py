class Solution:
    def isHappy(self, n: int) -> bool:
        seen_numbers = set()

        while n != 1:
            n = self._sumOfSquares(n)
            if n in seen_numbers:
                return False
            seen_numbers.add(n)

        return True

    def _sumOfSquares(self, n: int) -> int:
        result = 0
        while n != 0:
            result += (n % 10) ** 2
            n //= 10

        return result
