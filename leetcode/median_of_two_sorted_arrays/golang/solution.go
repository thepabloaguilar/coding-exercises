package golang

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	const largeNumber = 1<<31 - 1

	short, long := nums1, nums2
	if len(short) > len(long) {
		short, long = long, short
	}

	shortLen, longLen := len(short), len(long)
	halfLen := int((shortLen + longLen + 1) / 2)
	left, right := 0, shortLen

	for left <= right {
		cutShort := int((left + right) / 2)
		cutLong := halfLen - cutShort

		// Short partitions
		shortLeftMax := -largeNumber
		if cutShort > 0 {
			shortLeftMax = short[cutShort-1]
		}
		shortRightMin := largeNumber
		if cutShort < shortLen {
			shortRightMin = short[cutShort]
		}

		// Long partitions
		longLeftMax := -largeNumber
		if cutLong > 0 {
			longLeftMax = long[cutLong-1]
		}
		longRightMin := largeNumber
		if cutLong < longLen {
			longRightMin = long[cutLong]
		}

		isBalanced := shortLeftMax <= longRightMin && longLeftMax <= shortRightMin
		if isBalanced {
			if (shortLen+longLen)%2 == 0 {
				left := max(shortLeftMax, longLeftMax)
				right := min(shortRightMin, longRightMin)

				return float64(left+right) / 2
			}

			return float64(max(shortLeftMax, longLeftMax))
		} else if shortLeftMax > longRightMin {
			right = cutShort - 1
		} else {
			left = cutShort + 1
		}
	}

	return 0
}
