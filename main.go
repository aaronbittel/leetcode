package main

import (
	"fmt"
)

func main() {
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}
	// fmt.Println("Input:", gas, cost)
	// res := canCompleteCircuit(gas, cost)
	// sol := 3
	// fmt.Println("res:", res, "sol:", sol, res == sol)

	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println("Input", gas, cost)
	res := canCompleteCircuit(gas, cost)
	sol := -1
	fmt.Println("res:", res, "sol:", sol, res == sol)
}
