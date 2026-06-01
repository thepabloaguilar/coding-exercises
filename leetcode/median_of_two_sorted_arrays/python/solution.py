class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        shorterArr, longerArr = nums1, nums2
        if len(nums1) > len(nums2):
            shorterArr, longerArr = nums2, nums1

        lenShort, lenLong = len(shorterArr), len(longerArr)
        totalLen = lenShort + lenLong
        halfLen = (totalLen + 1) // 2

        searchLeft, searchRight = 0, lenShort
        while searchLeft <= searchRight:
            cutShort = (searchLeft + searchRight) // 2
            cutLong = halfLen - cutShort

            maxLeftShort = float('-inf') if cutShort == 0 else shorterArr[cutShort - 1]
            minRightShort = float('inf') if cutShort == lenShort else shorterArr[cutShort]

            maxLeftLong = float('-inf') if cutLong == 0 else longerArr[cutLong - 1]
            minRightLong = float('inf') if cutLong == lenLong else longerArr[cutLong]

            # Check if we have found the perfectly balanced partition
            isValidPartition = (maxLeftShort <= minRightLong) and (maxLeftLong <= minRightShort)
            if isValidPartition:
                if totalLen % 2 == 0:
                    leftMax = max(maxLeftShort, maxLeftLong)
                    rightMin = min(minRightShort, minRightLong)
                    return (leftMax + rightMin) / 2
                else:
                    return float(max(maxLeftShort, maxLeftLong))
            elif maxLeftShort > minRightLong:
                searchRight = cutShort - 1
            else:
                searchLeft = cutShort + 1
