package golang

func RemoveElement(nums []int, val int) int {
	itemToChange := len(nums) - 1

	for i := 0; i <= itemToChange; i++ {
		if nums[i] == val {
			nums[i] = nums[itemToChange]
			itemToChange--
			i--
		}
	}

	return itemToChange + 1
}
