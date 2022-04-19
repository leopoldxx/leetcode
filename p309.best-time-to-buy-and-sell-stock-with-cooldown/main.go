package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	} else if len(prices) == 2 {
		profit := prices[1] - prices[0]
		if profit < 0 {
			return 0
		}
		return profit
	}

	rest := 0
	buy := -prices[0]
	sell := -10000
	for i := 1; i < len(prices); i++ {
		preRest, preBuy := rest, buy
		rest = max(preRest, sell)
		buy = max(preRest-prices[i], preBuy)
		sell = preBuy + prices[i]
	}
	return max(rest, sell)
}
