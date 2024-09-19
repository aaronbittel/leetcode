package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(nums)
	fmt.Println(removeDuplicates2(nums))
	fmt.Println(nums)

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(nums)
	fmt.Println(removeDuplicates2(nums))
	fmt.Println(nums)
}
