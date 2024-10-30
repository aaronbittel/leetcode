package main

// 0ms
func moveZeroes(nums []int) {
	pos := 0
	for i, num := range nums {
		if num != 0 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos += 1
		}
	}
}

// 27ms
func moveZeroesFaster(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		for j := i; j < len(nums)-1; j++ {
			if nums[j+1] == 0 {
				break
			}
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}
}

// 44ms
func moveZeroesWorksButSlow(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		for j := i; j < len(nums)-1; j++ {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}
}
