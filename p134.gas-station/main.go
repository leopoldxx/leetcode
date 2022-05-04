package main

func canCompleteCircuit(gas []int, cost []int) int {
	if gas == nil || cost == nil {
		return -1
	}
	sum, total, index := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if sum < 0 {
			index = i + 1
			sum = 0
		}
	}
	if total < 0 {
		return -1
	}

	return index % len(gas)
}
