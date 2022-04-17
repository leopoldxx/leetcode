package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	total := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			total += profit
		}
	}
	return total
}
