package main

func canJump(nums []int) bool {
	steps := 1
	for i := 0; i < len(nums); i++ {
		steps--
		if nums[i] > steps {
			steps = nums[i]
		}

		if i+steps >= len(nums)-1 {
			return true
		}

		if steps == 0 {
			return false
		}
	}
	return false
}
