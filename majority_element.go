package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	maj := m[nums[0]]

	for _, n := range nums {
		m[n]++
		if m[n] > m[maj] {
			maj = n
		}
	}

	return maj
}

func majorityElementFastest(nums []int) int {
	result, count := 0, 0
	fmt.Println(nums)
	for _, num := range nums {
		fmt.Println(result, count)
		switch {
		case count == 0:
			result = num
			count = 1
		case num == result:
			count++
		default:
			count--
		}
	}
	fmt.Println(result, count)
	return result
}
