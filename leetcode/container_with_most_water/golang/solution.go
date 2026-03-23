package golang

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxContainer := min(height[left], height[right]) * (right - left)

	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

		maxContainer = max(
			maxContainer,
			min(height[left], height[right])*(right-left),
		)
	}

	return maxContainer
}
