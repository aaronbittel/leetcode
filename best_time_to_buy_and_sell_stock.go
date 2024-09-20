package main

func maxProfit(prices []int) int {
	bestProfit := 0
	lowestPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < lowestPrice {
			lowestPrice = prices[i]
			continue
		}
		pro := prices[i] - lowestPrice
		if pro > bestProfit {
			bestProfit = pro
		}
	}

	return bestProfit
}
