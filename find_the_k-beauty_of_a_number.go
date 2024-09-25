package main

import (
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	strNum := strconv.Itoa(num)
	count := 0

	for i := 0; i < len(strNum)-k+1; i++ {
		div, _ := strconv.Atoi(strNum[i : i+k])
		if div == 0 {
			continue
		}
		if num%div == 0 {
			count++
		}
	}
	return count
}
