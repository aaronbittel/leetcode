package main

func longestOnes(nums []int, k int) int {
	l := 0
	zeroCount := 0
	res := 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeroCount++
			for zeroCount > k {
				if nums[l] == 0 {
					zeroCount--
				}
				l++
			}
		}

		res = max(r-l+1, res)
	}

	return res
}
