package main

import (
	"fmt"
	"slices"
)

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	left, right := 0, 1

	// search for first none-zero value
	for height[left] == 0 {
		left++
		right++
	}

	volumn := 0

	for right < len(height) {
		if height[right] >= height[left] {
			limit := min(height[right], height[left])
			for i := left + 1; i < right; i++ {
				volumn += limit - height[i]
			}
			left = right
		}
		right++
	}

	if left >= len(height)-1 {
		return volumn
	}

	end := left
	right = len(height) - 1
	left = right - 1

	// Move from right to last left pointer
	for left >= end {
		if height[left] >= height[right] {
			limit := min(height[right], height[left])
			for i := right - 1; i > left; i-- {
				volumn += limit - height[i]
			}
			right = left
		}
		left--
	}

	return volumn
}

func trapV1(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	left, right := 0, 1

	// search for first none-zero value
	for height[left] == 0 {
		left++
		right++
	}

	volumn := 0
	lowest := 10_000

	for right < len(height) {
		if height[right] >= height[left] {
			limit := min(height[right], height[left])
			for i := left + 1; i < right; i++ {
				volumn += limit - height[i]
			}
			left = right
			right++
			lowest = 10_000
		} else {
			if lowest < height[right] {
				limit := height[right]
				for i := left + 1; i < right; i++ {
					amount := max(limit-height[i], 0)
					volumn += amount
					height[i] += amount
				}
				lowest = height[right]
				right++
				continue
			}
			lowest = min(lowest, height[right])
			right++
		}
	}

	return volumn
}

func printHeight(nums []int) {
	var (
		fullBlock = "█"
		maxVal    = slices.Max(nums)
		char      string
	)

	for row := maxVal; row >= 1; row-- {
		for _, n := range nums {
			if n >= row {
				char = fullBlock
			} else {
				char = " "
			}
			fmt.Print(fmt.Sprintf("%s ", char))
		}
		fmt.Println()
	}

	for _, n := range nums {
		fmt.Print(fmt.Sprintf("%d ", n))
	}
	fmt.Println()
}
