package golang

func CanCompleteCircuit(gas []int, cost []int) int {
	total := 0
	tank := 0
	startIdx := 0
	for idx := 0; idx < len(gas); idx++ {
		diff := gas[idx] - cost[idx]
		total += diff

		tank += diff
		if tank < 0 {
			tank = 0
			startIdx = idx + 1
		}
	}

	if total < 0 {
		return -1
	}

	return startIdx
}
