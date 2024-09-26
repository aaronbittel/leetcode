package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	sol := 7
	fmt.Println(prices)
	res := maxProfit2(prices)
	fmt.Println(res, sol, res == sol)
	fmt.Println("-------------")

	prices = []int{1, 2, 3, 4, 5}
	sol = 4
	fmt.Println(prices)
	res = maxProfit2(prices)
	fmt.Println(res, sol, res == sol)
	fmt.Println("-------------")

	prices = []int{7, 6, 4, 3, 1}
	sol = 0
	fmt.Println(prices)
	res = maxProfit2(prices)
	fmt.Println(res, sol, res == sol)
	fmt.Println("-------------")

}
