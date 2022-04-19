package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func max3(a, b, c int) int {
	if a <= c && b <= c {
		return c
	} else if a <= b && c <= b {
		return b
	}
	return a
}

func maxProfit(prices []int, fee int) int {
	rest := 0
	buy := -prices[0]
	sell := -1000000
	for i := 1; i < len(prices); i++ {
		preRest, preBuy := rest, buy
		rest = max(preRest, sell)
		buy = max3(preRest-prices[i], sell-prices[i], preBuy)
		sell = buy + prices[i] - fee
	}

	return max(rest, sell)
}
