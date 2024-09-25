package main

func findLHS(nums []int) int {
	maxLen := 0
	set := make(map[int]int, len(nums))

	for _, n := range nums {
		set[n]++
		if v, ok := set[n+1]; ok {
			maxLen = max(maxLen, set[n]+v)
		}

		if v, ok := set[n-1]; ok {
			maxLen = max(maxLen, set[n]+v)
		}
	}

	return maxLen
}

func findLHSV1(nums []int) int {
	set := map[int]int{}

	for _, n := range nums {
		set[n]++
	}

	maxLen := 0

	for k := range set {
		v, ok := set[k+1]
		if !ok {
			continue
		}
		maxLen = max(maxLen, set[k]+v)
	}

	return maxLen
}
