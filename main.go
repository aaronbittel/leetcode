package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 0, 6, 1, 5}
	sol := 3
	fmt.Println(nums)
	res := hIndex(nums)
	fmt.Println(res, sol, res == sol)
	fmt.Println("------------")

	nums = []int{1, 3, 1}
	sol = 1
	fmt.Println(nums)
	res = hIndex(nums)
	fmt.Println(res, sol, res == sol)
	fmt.Println("------------")

	nums = []int{9, 7, 6, 2, 1}
	sol = 3
	fmt.Println(nums)
	res = hIndex(nums)
	fmt.Println(res, sol, res == sol)
	fmt.Println("------------")
	//
	// nums = []int{3, 4, 3, 2, 5, 4, 3}
	// sol = 3
	// fmt.Println(nums)
	// res = jump(nums)
	// fmt.Println(res, sol, res == sol)
	// fmt.Println("------------")
}
