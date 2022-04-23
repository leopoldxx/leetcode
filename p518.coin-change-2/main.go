package main

func change(amount int, coins []int) int {
	preWays := make([]int, amount+1)
	currentWays := make([]int, amount+1)
	preWays[0] = 1

	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		for j := 0; j <= amount; j++ {
			if coin > j {
				currentWays[j] = preWays[j]
			} else {
				currentWays[j] = currentWays[j-coin] + preWays[j]
			}
		}
		preWays, currentWays = currentWays, preWays
	}
	return preWays[amount]
}
