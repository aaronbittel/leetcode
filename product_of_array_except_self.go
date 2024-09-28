package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prod := make([]int, n)

	prefix := 1
	suffixes := make([]int, n)
	suffixes[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffixes[i] = suffixes[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		prod[i] = suffixes[i] * prefix
		prefix *= nums[i]
	}

	return prod
}
