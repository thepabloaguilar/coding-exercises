class Solution:
    def merge(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        idx = (m + n) - 1
        m -= 1
        n -= 1
        while idx >= 0:
            if m < 0 <= n:
                nums1[idx] = nums2[n]
                n -= 1
            elif n < 0 <= m:
                nums1[idx] = nums1[m]
                m -= 1
            elif nums1[m] > nums2[n]:
                nums1[idx] = nums1[m]
                m -= 1
            else:
                nums1[idx] = nums2[n]
                n -= 1

            idx -=1
