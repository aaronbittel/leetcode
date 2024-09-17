package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	nums = []int{0}
	fmt.Println(nums)
	fmt.Println(canJump(nums))
	return

	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(nums)
	fmt.Println(canJump(nums))

	nums = []int{2, 0, 0}
	fmt.Println(nums)
	fmt.Println(canJump(nums))

	nums = []int{2, 5, 0, 0}
	fmt.Println(nums)
	fmt.Println(canJump(nums))
}
