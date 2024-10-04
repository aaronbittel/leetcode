package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0

	for left < right {
		lBar, rBar := height[left], height[right]
		rect := min(lBar, rBar) * (right - left)
		res = max(res, rect)
		if lBar <= rBar {
			left++
		} else {
			right--
		}
	}

	return res
}

func maxAreaBruteForce(height []int) int {
	res := 0

	for l := 0; l < len(height); l++ {
		for r := l + 1; r < len(height); r++ {
			rect := min(height[l], height[r]) * (r - l)
			res = max(res, rect)
		}
	}

	return res
}
