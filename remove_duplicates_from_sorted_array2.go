package main

func removeDuplicates2(nums []int) int {
	cur, last, count := 1, nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			if count == 1 {
				nums[i], nums[cur] = nums[cur], nums[i]
				cur++
				count++
			}
			continue
		}
		last = nums[i]
		nums[i], nums[cur] = nums[cur], nums[i]
		cur++
		count = 1
	}
	return cur
}
