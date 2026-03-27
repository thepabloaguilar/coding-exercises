class NumArray:
    def __init__(self, nums: list[int]):
        self._prefix_sum = []

        s = 0
        for num in nums:
            s += num
            self._prefix_sum.append(s)

    def sumRange(self, left: int, right: int) -> int:
        if left == 0:
            return self._prefix_sum[right]
        return self._prefix_sum[right] - self._prefix_sum[left - 1]
