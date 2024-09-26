package main

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	diff := prices[1] - prices[0]
	var profit int
	if diff > 0 {
		profit = diff
	}
	for i := 1; i < len(prices)-1; i++ {
		diff = prices[i+1] - (diff + prices[i-1])
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}

func maxProfit2V1(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		diff := prices[i+1] - prices[i]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}
