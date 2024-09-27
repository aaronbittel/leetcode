package main

func jump(nums []int) int {
	jumps, l, r := 0, 0, 0

	for r < len(nums)-1 {
		m := -1
		for i := l; i <= r; i++ {
			m = max(m, i+nums[i])
		}
		l = r + 1
		r = m
		jumps++
	}
	return jumps
}
