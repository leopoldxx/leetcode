package main

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	pre_1 := cost[1]
	pre_2 := cost[0]
	for i := 2; i < len(cost); i++ {
		res := min(cost[i]+pre_1, cost[i]+pre_2)
		pre_2, pre_1 = pre_1, res
	}
	return min(pre_2, pre_1)
}
