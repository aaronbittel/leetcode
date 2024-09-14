package main

func rotate(nums []int, k int) {
	l := len(nums)
	_ = l
}

func rotateWithHelper(nums []int, k int) {
	l := len(nums)
	help := make([]int, l)
	for i := 0; i < l; i++ {
		help[(i+k)%l] = nums[i]
	}
	for i, n := range help {
		nums[i] = n
	}
}
