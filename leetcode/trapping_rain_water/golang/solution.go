package golang

func Trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	waterAmount := 0

	for left < right {
		if leftMax <= rightMax {
			left++
			leftMax = max(leftMax, height[left])
			waterAmount += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			waterAmount += rightMax - height[right]
		}
	}

	return waterAmount
}
