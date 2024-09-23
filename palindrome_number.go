package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var mirror int
	y := x
	for y != 0 {
		mirror = mirror*10 + y%10
		y /= 10
	}

	return mirror == x
}

func isPalindromeArray(x int) bool {
	if x < 0 {
		return false
	}

	var nums []int
	for x != 0 {
		nums = append(nums, x%10)
		x /= 10
	}

	length := len(nums)

	for i := 0; i < length/2; i++ {
		if nums[i] != nums[length-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeConvertToString(x int) bool {
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	l := len(xStr)
	for i := 0; i < l/2; i++ {
		if xStr[i] != xStr[l-1-i] {
			return false
		}
	}
	return true
}
