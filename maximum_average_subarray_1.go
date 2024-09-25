package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	curSum := 0
	maxSum := math.Inf(-1)

	for i := range k {
		curSum += nums[i]
	}
	maxSum = float64(curSum)

	for i := k; i < len(nums); i++ {
		curSum += nums[i]
		curSum -= nums[i-k]
		maxSum = max(maxSum, float64(curSum))
	}

	return maxSum / float64(k)
}
