package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numMap := make(map[int]struct{})
	longest := 1

	for _, n := range nums {
		numMap[n] = struct{}{}
	}

	for start := range numMap {
		if _, ok := numMap[start-1]; ok {
			continue
		}

		// start is a valid starting point for a sequence
		length := 1
		for {
			if _, ok := numMap[start+length]; ok {
				length++
				continue
			}
			break
		}

		longest = max(longest, length)
	}

	return longest
}
