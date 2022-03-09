package main

func maxProfit(prices []int) int {
	maxProfit := 0
	minBuy := prices[0]
	for _, price := range prices {
		profit := price - minBuy
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minBuy {
			minBuy = price
		}
	}
	return maxProfit
}
