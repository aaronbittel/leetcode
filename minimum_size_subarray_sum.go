package main

func minSubArrayLen(target int, nums []int) int {
	left, curLen := 0, 0
	minLen := len(nums) + 1

	for right := range len(nums) {
		curLen += nums[right]
		for curLen >= target {
			minLen = min(minLen, right-left+1)
			curLen -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}

func minSubArrayLenV1(target int, nums []int) int {
	left, right := 0, 0
	minLen := len(nums) + 1
	curLen := nums[left]

	for right < len(nums) {
		if curLen < target {
			right++
			if right >= len(nums) {
				break
			}
			curLen += nums[right]
			continue
		}
		minLen = min(minLen, right-left+1)
		curLen -= nums[left]
		left++
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
