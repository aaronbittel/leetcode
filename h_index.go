package main

func hIndex(citations []int) int {
	bubblesort(citations)

	h := 1
	for {
		var i int
		for i = 0; i < len(citations); i++ {
			if citations[i] >= h {
				break
			}
		}
		if len(citations)-i >= h {
			h++
			continue
		}
		break
	}

	return h - 1
}

func bubblesort(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
