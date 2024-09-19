package main

func removeDuplicates(nums []int) int {
	cur, last := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			continue
		}
		last = nums[i]
		nums[i], nums[cur] = nums[cur], nums[i]
		cur++
	}
	return cur
}
